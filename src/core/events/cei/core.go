package cei

import (
	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

func BasicCEIEvents(customers []mapper.Customer, corporateActions map[string][]mapper.CorporateAction) {
	CEITransaction := repository.GetCEITransaction(customers)
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
				CEITransactionPersisterObject = append(CEITransactionPersisterObject, ApplyBasicCorporateAction(transaction, corporate_action))
				continue
			}
		}
	}

	repository.UpdateCEITransaction(CEITransactionPersisterObject)
}

func ProceedsCEIEvents(corporateActions map[string][]mapper.CorporateAction, customers []mapper.Customer, symbols []mapper.Symbol) {
	CEITransactions := repository.GetAllCEITransactions(customers)
	CEIProceedPersisterObject := []mapper.CEIProceeds{}
	for _, customer := range customers {

		for _, symbol := range symbols {
			CEIProceedPersisterObject = append(CEIProceedPersisterObject, ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, CEITransactions, corporateActions)...)

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
