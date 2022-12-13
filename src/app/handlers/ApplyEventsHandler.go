package handlers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/core/events/oms"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/interfaces"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/repository"
)

type ApplyEventsHandler struct {
	_corporateActionsRepository interfaces.ICorporateActionRepository
	_customerRepository         interfaces.ICustomerRepository
}

var CorporateActionsAsc map[string][]mapper.CorporateAction
var CorporateActionsDesc map[string][]mapper.CorporateAction
var hashLogID string

func (t ApplyEventsHandler) NewApplyEventsHandler() error {
	return ApplyEventsHandler{
		_corporateActionsRepository: repository.NewCorporateActionRepository(),
		_customerRepository:         repository.NewCustomerRepository(),
	}.handler()
}

func (t ApplyEventsHandler) handler() error {
	start := time.Now()

	getCorporateActions, err := t._corporateActionsRepository.GetAll()
	if err != nil {
		return err
	}

	t.ApplyEvents(constants.AllCustomers)
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

func (t ApplyEventsHandler) ApplyEvents(customerCode string) {
	id := uuid.New()
	hashLogID = id.String()

	corporateActions, err := t._corporateActionsRepository.GetAll()
	if err != nil {

	}

	OMSCustomers := []dtos.Customer{}
	if customerCode == constants.AllCustomers {
		OMSCustomers, err = t._customerRepository.GetAll()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
				"Error":  err,
			}).Error("erro consultando GetOMSCustomers")
		}
	} else {
		OMSCustomers = []dtos.Customer{
			dtos.Customer{
				CustomerCode: customerCode,
				CreatedAT:    time.Now().String(),
			},
		}
	}
	oms.BasicOMSEvents(OMSCustomers, corporateActions.Map)
	//err := applyAllEventsOMS(customerCode)
	//if err != nil {
	//	return
	//}

	walletConnector := repository.NewWalletConnector()
	walletConnector.ResyncAVGOMS()

}

func applyAllEventsOMS(customerCode string) error {
	//var wg sync.WaitGroup

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

	// wg.Add(2)
	// go func() {
	// 	defer wg.Done()
	oms.BasicOMSEvents(OMSCustomers, CorporateActionsDesc)
	//}()

	// go func() {
	// 	defer wg.Done()
	// 	doProceedsOMSEvents(false, OMSCustomers)
	// }()

	// wg.Wait()

	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos OMS")

	return nil
}

// func doProceedsOMSEvents(isStateLess bool, OMSCustomers []mapper.Customer) {
// 	OMSSymbols, err := repository.GetOMSSymbols(OMSCustomers)
// 	if err != nil {
// 		logrus.WithFields(logrus.Fields{
// 			"HashID": hashLogID,
// 			"Error":  err,
// 		}).Error("erro consultando OMSSymbols, segue o jogo")
// 		return
// 	}

// 	oms.ProceedsOMSEvents(CorporateActionsAsc, OMSCustomers, OMSSymbols, isStateLess)
// }
