package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/cei"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/sirupsen/logrus"
)

func Run() {
	start := time.Now()
	ApplyEvents(constants.AllCustomers)
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

var CorporateActionsAsc map[string][]mapper.CorporateAction
var CorporateActionsDesc map[string][]mapper.CorporateAction
var hashLogID string

func ApplyEvents(customerCode string) {
	id := uuid.New()
	hashLogID = id.String()

	CorporateActionsAsc = repository.GetCorporateActions("asc")
	CorporateActionsDesc = repository.GetCorporateActions("desc")

	err := applyAllEventsInvestor(customerCode)
	if err != nil {
		return
	}

	err = applyAllEventsManual(customerCode)
	if err != nil {
		return
	}
	CorporateActionsAsc = map[string][]mapper.CorporateAction{}
	CorporateActionsDesc = map[string][]mapper.CorporateAction{}
	walletConnector := repository.NewWalletConnector()
	walletConnector.ResyncAVGInvestor()
	walletConnector.ResyncAVGManual()
}

func ApplyEventsAfterInvestorSync(customerCode string) error {
	options.OPTIONS.Load()
	CorporateActionsAsc = repository.GetCorporateActions("asc")
	CorporateActionsDesc = repository.GetCorporateActions("desc")

	Customer := mapper.Customer{
		CustomerCode: customerCode,
		CreatedAT:    time.Now().String(),
	}

	CEICustomers := []mapper.Customer{Customer}

	cei.BasicCEIEvents(CEICustomers, CorporateActionsDesc, true)
	doProceedsCEIEventsServerless(true, CEICustomers)

	return nil
}

func applyAllEventsInvestor(customerCode string) error {
	var wg sync.WaitGroup
	customers := []mapper.Customer{}
	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Inicia Eventos Portal do Investidor")

	if customerCode == constants.AllCustomers {
		GetCEICustomers, err := repository.GetCEICustomers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
				"Error":  err,
			}).Error("erro consultando GetCEICustomers")
			return err
		}

		if GetCEICustomers == nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
			}).Info("GetCEICustomers not found")
			return nil
		}
		customers = GetCEICustomers
	} else {
		Customer := mapper.Customer{
			CustomerCode: customerCode,
			CreatedAT:    time.Now().String(),
		}

		customers = []mapper.Customer{Customer}
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		cei.BasicCEIEvents(customers, CorporateActionsDesc, false)
	}()
	go func() {
		defer wg.Done()
		doProceedsCEIEvents(false, customers)
	}()

	wg.Wait()

	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos Portal do Investidor")
	customers = []mapper.Customer{}
	return nil
}

func applyAllEventsManual(customerCode string) error {
	customers := []mapper.Customer{}
	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Inicia Eventos Manuais")

	if customerCode == constants.AllCustomers {
		GetManualCustomers, err := repository.GetManualCustomers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
				"Error":  err,
			}).Error("erro consultando GetManualCustomers")
			return err
		}

		if GetManualCustomers == nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
			}).Info("GetManualCustomers not found")
			return nil
		}
		customers = GetManualCustomers
	} else {
		Customer := mapper.Customer{
			CustomerCode: customerCode,
			CreatedAT:    time.Now().String(),
		}

		customers = []mapper.Customer{Customer}
	}

	manual.BasicManualEvents(customers, CorporateActionsDesc)
	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos Manuais")
	customers = []mapper.Customer{}
	return nil
}

func doProceedsCEIEvents(isStateLess bool, CEICustomers []mapper.Customer) {
	CEISymbols, err := repository.GetCEISymbols(CEICustomers)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"HashID": hashLogID,
			"Error":  err,
		}).Error("erro consultando CEISymbols")
		return
	}

	cei.ProceedsCEIEvents(CorporateActionsAsc, CEICustomers, CEISymbols, isStateLess)
}

func doProceedsCEIEventsServerless(isStateLess bool, CEICustomers []mapper.Customer) {
	CEISymbols, err := repository.GetCEISymbols(CEICustomers)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"HashID": hashLogID,
			"Error":  err,
		}).Error("erro consultando CEISymbols, segue o jogo")
		return
	}
	cei.ProceedsCEIEvents(CorporateActionsAsc, CEICustomers, CEISymbols, isStateLess)
}
