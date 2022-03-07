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
	Customers := repository.GetCustomers()
	Symbols := repository.GetSymbols()

	wg.Add(5)
	go doBasicOMSEvents(CorporateActionsDesc)
	go doBasicManualEvents(CorporateActionsDesc)
	go doBasicCEIEvents(CorporateActionsDesc)
	go doProceedsOMSEvents(CorporateActionsAsc, Customers, Symbols)
	go doProceedsCEIEvents(CorporateActionsAsc, Customers, Symbols)
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

			// Se a data de EventDate for maior, significa que eventos corporativos com datas inferiores, já foram aplicados (Eventos com ano 2001 sao eventos com data default e devem ser consideraros pois é a primeira vez)
			if corporate_action.ComDate.After(transaction.EventDate) && transaction.EventDate.Year() > 2001 {
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

			// Se a data de EventDate for maior, significa que eventos corporativos com datas inferiores, já foram aplicados (Eventos com ano 2001 sao eventos com data default e devem ser consideraros pois é a primeira vez)
			if corporate_action.ComDate.After(transaction.EventDate) && transaction.EventDate.Year() > 2001 {
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

			// Se a data de EventDate for maior, significa que eventos corporativos com datas inferiores, já foram aplicados (Eventos com ano 2001 sao eventos com data default e devem ser consideraros pois é a primeira vez)
			if corporate_action.ComDate.After(transaction.EventDate) && transaction.EventDate.Year() > 2001 {
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

func doProceedsOMSEvents(corporateActions map[string][]mapper.CorporateAction, customers []mapper.Customer, symbols []mapper.Symbol) {
	defer wg.Done()
	OMSTransactions := repository.GetAllOMSTransactions()
	OMSProceedPersisterObject := []mapper.OMSProceeds{}
	for _, customer := range customers {

		for _, symbol := range symbols {
			OMSProceedPersisterObject = append(OMSProceedPersisterObject, oms.ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, OMSTransactions, corporateActions)...)

		}
	}

	if len(OMSProceedPersisterObject) > 0 {
		repository.InsertOMSProceeds(OMSProceedPersisterObject)
	}

	ManualTransactions := []mapper.ManualTransaction{}
	for _, proceed := range OMSProceedPersisterObject {
		if proceed.Event == constants.Bonus {
			ManualTransaction := mapper.ManualTransaction{}
			ManualTransactions = append(ManualTransactions, manual.ApplyInheritedBonusActionOMS(ManualTransaction, proceed))
		}
	}

	if len(ManualTransactions) > 0 {
		repository.InsertManualTransaction(ManualTransactions)
	}
}

func doProceedsCEIEvents(corporateActions map[string][]mapper.CorporateAction, customers []mapper.Customer, symbols []mapper.Symbol) {
	defer wg.Done()
	CEITransactions := repository.GetAllCEITransactions()
	CEIProceedPersisterObject := []mapper.CEIProceeds{}
	for _, customer := range customers {

		for _, symbol := range symbols {
			CEIProceedPersisterObject = append(CEIProceedPersisterObject, cei.ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, CEITransactions, corporateActions)...)

		}
	}

	if len(CEIProceedPersisterObject) > 0 {
		repository.InsertCEIProceeds(CEIProceedPersisterObject)
	}

	ManualTransactions := []mapper.ManualTransaction{}
	for _, proceed := range CEIProceedPersisterObject {
		if proceed.Event == constants.Bonus {
			ManualTransaction := mapper.ManualTransaction{}
			ManualTransactions = append(ManualTransactions, manual.ApplyInheritedBonusActionCEI(ManualTransaction, proceed))
		}
	}

	if len(ManualTransactions) > 0 {
		repository.InsertManualTransaction(ManualTransactions)
	}
}
