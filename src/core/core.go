package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/cei"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

var wg sync.WaitGroup

func Run() {
	start := time.Now()
	doApplyBasicEvents()
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

func doApplyBasicEvents() {
	// CorporateActionsAsc := repository.GetCorporateActions("asc")
	CorporateActionsDesc := repository.GetCorporateActions("desc")
	wg.Add(1)
	// go doBasicOMSEvents(CorporateActionsDesc)
	go doBasicManualEvents(CorporateActionsDesc)
	// go doBasicCEIEvents(CorporateActionsDesc)
	// go doProceedsOMSEvents(CorporateActionsAsc)
	wg.Wait()
}

func doBasicOMSEvents(corporateActions map[string][]mapper.CorporateAction) {
	defer wg.Done()
	OMSTransaction := repository.GetOMSTransaction()
	OMSTransactionPersisterObject := []mapper.OMSTransaction{}

	for _, transaction := range OMSTransaction {

		for _, corporate_action := range corporateActions[transaction.Symbol] {

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if corporate_action.IsBasic() {
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

func doBasicManualEvents(corporateActions map[string][]mapper.CorporateAction) {
	defer wg.Done()
	ManualTransaction := repository.GetManualTransaction()
	ManualTransactionPersisterObject := []mapper.ManualTransaction{}

	for _, transaction := range ManualTransaction {

		for _, corporate_action := range corporateActions[transaction.Symbol] {

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if corporate_action.IsBasic() {
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

func doBasicCEIEvents(corporateActions map[string][]mapper.CorporateAction) {
	defer wg.Done()
	CEITransaction := repository.GetCEITransaction()
	CEITransactionPersisterObject := []mapper.CEITransaction{}

	for _, transaction := range CEITransaction {

		for _, corporate_action := range corporateActions[transaction.Symbol] {

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if corporate_action.IsBasic() {
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

func doProceedsOMSEvents(corporateActions map[string][]mapper.CorporateAction) {
	defer wg.Done()
	OMSTransactions := repository.GetAllOMSTransactions()
	OMSCustomers := repository.GetCustomers()
	OMSSymbols := repository.GetSymbols()
	OMSProceedPersisterObject := []mapper.OMSProceeds{}
	for _, customer := range OMSCustomers {

		for _, symbol := range OMSSymbols {
			OMSProceedPersisterObject = append(OMSProceedPersisterObject, oms.ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, OMSTransactions, CorporateActions)...)

		}
	}
	// TODO - Validar bem os dados de proventos.
	// O ideal seria fazer os Inserts com base em Symbols de um determinado customer.
	// Mas fazer 1 select por customer é loucura. Demora demais.
	// Validei com uns 20 symbols e parece estar tudo bem, mas vou manter esse todo para validar um pouco mais
	repository.InsertOMSProceeds(OMSProceedPersisterObject)

	ManualTransactions := []mapper.ManualTransaction{}
	for _, proceed := range OMSProceedPersisterObject {
		if proceed.Event == constants.Bonus && proceed.Quantity > 0 {
			ManualTransaction := mapper.ManualTransaction{}

			ManualTransaction.CustomerCode = proceed.CustomerCode
			ManualTransaction.BrokerID = proceed.BrokerID
			ManualTransaction.InvestmentType = constants.BonusInvestmentType
			ManualTransaction.Symbol = proceed.Symbol
			ManualTransaction.Quantity = proceed.Quantity
			ManualTransaction.Price = constants.MinimalPrice
			ManualTransaction.Amount = ManualTransaction.Quantity * ManualTransaction.Price
			ManualTransaction.Side = constants.Purchase
			ManualTransaction.TradeDate = proceed.Date // TODO com_date ou initial_date do evento ?
			ManualTransaction.SourceType = constants.ManualSourceType
			ManualTransaction.EventDate = proceed.Date
			ManualTransaction.EventName = proceed.Event

			ManualTransactions = append(ManualTransactions, ManualTransaction)
		}
	}

	repository.InsertManualTransaction(ManualTransactions)

}
