package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/core/events/cei"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

var wg sync.WaitGroup

func Run() {
	start := time.Now()
	doApplyBasicEvents()
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

func doApplyBasicEvents() {
	CorporateActions := repository.GetCorporateActions()
	wg.Add(3)
	go doBasicOMSEvents(CorporateActions)
	go doBasicManualEvents(CorporateActions)
	go doBasicCEIEvents(CorporateActions)
	wg.Wait()
}

func doBasicOMSEvents(CorporateActions map[string][]mapper.CorporateAction) {
	defer wg.Done()
	OMSTransaction := repository.GetOMSTransaction()
	OMSTransactionPersisterObject := []mapper.OMSTransaction{}

	for _, transaction := range OMSTransaction {

		for _, corporate_action := range CorporateActions[transaction.Symbol] {

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if utils.Contains([]string{singleton.New().Grouping, singleton.New().Unfolding, singleton.New().Update}, corporate_action.Description) {
				transaction.EventName = corporate_action.Description
				transaction.PostEventSymbol = corporate_action.TargetTicker
				transaction.EventFactor = corporate_action.CalculatedFactor
				transaction.EventDate = corporate_action.ComDate
				OMSTransactionPersisterObject = append(OMSTransactionPersisterObject, oms.ApplyBasicCorporateAction(transaction))
				continue
			}
		}
	}

	repository.UpdateOMSTransaction(OMSTransactionPersisterObject)
}

func doBasicManualEvents(CorporateActions map[string][]mapper.CorporateAction) {
	defer wg.Done()
	ManualTransaction := repository.GetManualTransaction()
	ManualTransactionPersisterObject := []mapper.ManualTransaction{}

	for _, transaction := range ManualTransaction {

		for _, corporate_action := range CorporateActions[transaction.Symbol] {

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if utils.Contains([]string{singleton.New().Grouping, singleton.New().Unfolding, singleton.New().Update}, corporate_action.Description) {
				transaction.EventName = corporate_action.Description
				transaction.PostEventSymbol = corporate_action.TargetTicker
				transaction.EventFactor = corporate_action.CalculatedFactor
				transaction.EventDate = corporate_action.ComDate
				ManualTransactionPersisterObject = append(ManualTransactionPersisterObject, manual.ApplyBasicCorporateAction(transaction))
				continue
			}
		}
	}

	repository.UpdateManualTransaction(ManualTransactionPersisterObject)
}

func doBasicCEIEvents(CorporateActions map[string][]mapper.CorporateAction) {
	defer wg.Done()
	CEITransaction := repository.GetCEITransaction()
	CEITransactionPersisterObject := []mapper.CEITransaction{}

	for _, transaction := range CEITransaction {

		for _, corporate_action := range CorporateActions[transaction.Symbol] {

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if utils.Contains([]string{singleton.New().Grouping, singleton.New().Unfolding, singleton.New().Update}, corporate_action.Description) {
				transaction.EventName = corporate_action.Description
				transaction.PostEventSymbol = corporate_action.TargetTicker
				transaction.EventFactor = corporate_action.CalculatedFactor
				transaction.EventDate = corporate_action.ComDate
				CEITransactionPersisterObject = append(CEITransactionPersisterObject, cei.ApplyBasicCorporateAction(transaction))
				continue
			}
		}
	}

	repository.UpdateCEITransaction(CEITransactionPersisterObject)
}
