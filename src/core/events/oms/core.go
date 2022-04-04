package oms

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
)

func BasicOMSEvents(customers []mapper.Customer, corporateActions map[string][]mapper.CorporateAction) {
	OMSTransaction := repository.GetOMSTransaction(customers)
	OMSTransactionPersisterObject := []mapper.OMSTransaction{}

	for _, transaction := range OMSTransaction {

		for _, corporate_action := range corporateActions[transaction.Symbol] {

			// Se a data de InitialDate for maior, significa que eu não precios aplicar este evento nesta transação
			if transaction.TradeDate.After(corporate_action.InitialDate) && !transaction.TradeDate.Equal(corporate_action.InitialDate) {
				continue
			}

			// Se a data de EventDate for maior, significa que eventos corporativos com datas inferiores, já foram aplicados (Eventos com ano 2001 sao eventos com data default e devem ser consideraros pois é a primeira vez)
			if corporate_action.InitialDate.After(transaction.EventDate) && transaction.EventDate.Year() > 2001 {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if corporate_action.IsBasic() {
				OMSTransactionPersisterObject = append(OMSTransactionPersisterObject, ApplyBasicCorporateAction(transaction, corporate_action))
				continue
			}
		}
	}

	repository.UpdateOMSTransaction(OMSTransactionPersisterObject)
}

func ProceedsOMSEvents(corporateActions map[string][]mapper.CorporateAction, customers []mapper.Customer, symbols []mapper.Symbol) {
	OMSTransactions := repository.GetAllOMSTransactions(customers)
	OMSProceedPersisterObject := []mapper.OMSProceeds{}
	for _, customer := range customers {
		//talez aqui
		for _, symbol := range symbols {
			OMSProceedPersisterObject = append(OMSProceedPersisterObject, ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, OMSTransactions, corporateActions)...)

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
