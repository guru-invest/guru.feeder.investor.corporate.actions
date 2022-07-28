package core

import (
	"fmt"
	"sync"
	"time"

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

func ApplyEvents(customerCode string) {

	CorporateActionsAsc = repository.GetCorporateActions("asc")
	CorporateActionsDesc = repository.GetCorporateActions("desc")

	err := applyAllEventsOMS(customerCode)
	if err != nil {
		return
	}

	err = applyAllEventsManual(customerCode)
	if err != nil {
		return
	}

	err = applyAllEventsInvestor(customerCode)
	if err != nil {
		return
	}

	go repository.NewWalletConnector().ResyncAveragePrice()
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
	CEICustomers := []mapper.Customer{}
	logrus.Info("Inicia Eventos Portal do Investidor")
	if customerCode == constants.AllCustomers {
		GetCEICustomers, err := repository.GetCEICustomers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"GetCEICustomers": CEICustomers,
				"Error":           err,
			}).Error("erro consultando GetCEICustomers")
			return err
		}

		if CEICustomers == nil {
			logrus.WithFields(logrus.Fields{
				"GetCEICustomers": CEICustomers,
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
	logrus.Info("Finaliza Eventos Portal do Investidor")
	return nil
}

func applyAllEventsManual(customerCode string) error {
	ManualCustomers := []mapper.Customer{}
	logrus.Info("Inicia Eventos Manuais")
	if customerCode == constants.AllCustomers {
		GetManualCustomers, err := repository.GetManualCustomers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"GetManualCustomers": GetManualCustomers,
				"Error":              err,
			}).Error("erro consultando GetManualCustomers")
			return err
		}
		logrus.WithFields(logrus.Fields{
			"Total": len(GetManualCustomers),
		}).Info("Pega total dos manuais")
		
		if ManualCustomers == nil {
			logrus.WithFields(logrus.Fields{
				"GetManualCustomers": GetManualCustomers,
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
	logrus.Info("Finaliza Eventos Manuais")
	return nil
}

func applyAllEventsOMS(customerCode string) error {
	var wg sync.WaitGroup

	logrus.Info("Inicia Eventos OMS")

	OMSCustomers := []mapper.Customer{}
	if customerCode == constants.AllCustomers {
		GetOMSCustomers, err := repository.GetOMSCustomers()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"GetOMSCustomers": GetOMSCustomers,
				"Error":           err,
			}).Error("erro consultando GetOMSCustomers")
			return err
		}
		if OMSCustomers == nil {
			logrus.WithFields(logrus.Fields{
				"GetOMSCustomers": GetOMSCustomers,
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
	logrus.Info("Finaliza Eventos OMS")

	return nil
}

func doProceedsOMSEvents(isStateLess bool, OMSCustomers []mapper.Customer) {
	OMSSymbols, err := repository.GetOMSSymbols(OMSCustomers)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"OMSSymbols": OMSSymbols,
			"Error":      err,
		}).Error("erro consultando OMSSymbols, segue o jogo")
		return
	}

	oms.ProceedsOMSEvents(CorporateActionsAsc, OMSCustomers, OMSSymbols, isStateLess)
}

func doProceedsCEIEvents(isStateLess bool, CEICustomers []mapper.Customer) {
	CEISymbols, err := repository.GetCEISymbols(CEICustomers)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"CEISymbols": CEISymbols,
			"Error":      err,
		}).Error("erro consultando CEISymbols")
		return
	}
	cei.ProceedsCEIEvents(CorporateActionsAsc, CEICustomers, CEISymbols, isStateLess)
}

func doProceedsCEIEventsServerless(isStateLess bool, CEICustomers []mapper.Customer) {
	CEISymbols, err := repository.GetCEISymbols(CEICustomers)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"CEISymbols": CEISymbols,
			"Error":      err,
		}).Error("erro consultando CEISymbols, segue o jogo")
		return
	}
	cei.ProceedsCEIEvents(CorporateActionsAsc, CEICustomers, CEISymbols, isStateLess)
}
