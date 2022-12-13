package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/oms"
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

	err := applyAllEventsOMS(customerCode)
	if err != nil {
		return
	}

	walletConnector := repository.NewWalletConnector()
	walletConnector.ResyncAVGOMS()

}

func applyAllEventsOMS(customerCode string) error {

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

	oms.BasicOMSEvents(OMSCustomers, CorporateActionsDesc)

	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos OMS")

	return nil
}
