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
	CorporateActionsAsc := repository.GetCorporateActions("asc")
	CorporateActionsDesc := repository.GetCorporateActions("desc")
	wg.Add(4)
	go doBasicOMSEvents(CorporateActionsDesc)
	go doBasicManualEvents(CorporateActionsDesc)
	go doBasicCEIEvents(CorporateActionsDesc)
	go doProceedsOMSEvents(CorporateActionsAsc)
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
				OMSTransactionPersisterObject = append(OMSTransactionPersisterObject, oms.ApplyBasicCorporateAction(transaction, corporate_action))
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
				ManualTransactionPersisterObject = append(ManualTransactionPersisterObject, manual.ApplyBasicCorporateAction(transaction, corporate_action))
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
				CEITransactionPersisterObject = append(CEITransactionPersisterObject, cei.ApplyBasicCorporateAction(transaction, corporate_action))
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
			OMSProceedPersisterObject = append(OMSProceedPersisterObject, oms.ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, OMSTransactions, corporateActions)...)

		}
	}

	repository.InsertOMSProceeds(OMSProceedPersisterObject)

	ManualTransactions := []mapper.ManualTransaction{}
	for _, proceed := range OMSProceedPersisterObject {
		if proceed.Event == constants.Bonus {
			ManualTransaction := mapper.ManualTransaction{}
			ManualTransactions = append(ManualTransactions, manual.ApplyInheritedBonusAction(ManualTransaction, proceed))
		}
	}

	repository.InsertManualTransaction(ManualTransactions)

}
