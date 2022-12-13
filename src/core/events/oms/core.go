package oms

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/repository"
)

func BasicOMSEvents(customers []dtos.Customer, corporateActions domain.CorporateActionList) {
	OMSTransaction := repository.GetOMSTransaction(customers)
	OMSTransactionPersisterObject := []mapper.Transaction{}
	for _, transaction := range OMSTransaction {

		for _, corporate_action := range corporateActions.Map[transaction.Symbol] {

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
