package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/cei"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/oms"
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

	teste, _ := repository.GetCorporateActions("asc")
	teste2, _ := repository.GetCorporateActions("desc")
	fmt.Println(teste, teste2)
	//CorporateActionsDesc = repository.GetCorporateActions("desc")

	err := applyAllEventsOMS(customerCode)
	if err != nil {
		return
	}

	err = applyAllEventsInvestor(customerCode)
	if err != nil {
		return
	}

	err = applyAllEventsManual(customerCode)
	if err != nil {
		return
	}

	walletConnector := repository.NewWalletConnector()
	walletConnector.ResyncAVGOMS()
	walletConnector.ResyncAVGInvestor()
	walletConnector.ResyncAVGManual()
}

func ApplyEventsAfterInvestorSync(customerCode string) error {
	options.OPTIONS.Load()
	// CorporateActionsAsc = repository.GetCorporateActions("asc")
	// CorporateActionsDesc = repository.GetCorporateActions("desc")

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
	CEICustomers := []mapper.Customer{}
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
		CEICustomers = GetCEICustomers
	} else {
		Customer := mapper.Customer{
			CustomerCode: customerCode,
			CreatedAT:    time.Now().String(),
		}

		CEICustomers = []mapper.Customer{Customer}
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		cei.BasicCEIEvents(CEICustomers, CorporateActionsDesc, false)
	}()
	go func() {
		defer wg.Done()
		doProceedsCEIEvents(false, CEICustomers)
	}()

	wg.Wait()

	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos Portal do Investidor")
	return nil
}

func applyAllEventsManual(customerCode string) error {
	ManualCustomers := []mapper.Customer{}
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
		ManualCustomers = GetManualCustomers
	} else {
		Customer := mapper.Customer{
			CustomerCode: customerCode,
			CreatedAT:    time.Now().String(),
		}

		ManualCustomers = []mapper.Customer{Customer}
	}

	manual.BasicManualEvents(ManualCustomers, CorporateActionsDesc)
	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos Manuais")
	return nil
}

func applyAllEventsOMS(customerCode string) error {
	var wg sync.WaitGroup

	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Inicia Eventos OMS")

	OMSCustomers := []mapper.Customer{}
	if customerCode == constants.AllCustomers {
		GetOMSCustomers, err := repository.GetOMSCustomers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
				"Error":  err,
			}).Error("erro consultando GetOMSCustomers")
			return err
		}
		if OMSCustomers == nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
			}).Info("GetOMSCustomers not found")
			return nil
		}
		OMSCustomers = GetOMSCustomers
	} else {
		Customer := mapper.Customer{
			CustomerCode: customerCode,
			CreatedAT:    time.Now().String(),
		}

		OMSCustomers = []mapper.Customer{Customer}
	}

	wg.Add(2)
	go func() {
		defer wg.Done()
		oms.BasicOMSEvents(OMSCustomers, CorporateActionsDesc)
	}()

	go func() {
		defer wg.Done()
		doProceedsOMSEvents(false, OMSCustomers)
	}()

	wg.Wait()

	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos OMS")

	return nil
}

func doProceedsOMSEvents(isStateLess bool, OMSCustomers []mapper.Customer) {
	OMSSymbols, err := repository.GetOMSSymbols(OMSCustomers)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"HashID": hashLogID,
			"Error":  err,
		}).Error("erro consultando OMSSymbols, segue o jogo")
		return
	}

	oms.ProceedsOMSEvents(CorporateActionsAsc, OMSCustomers, OMSSymbols, isStateLess)
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
