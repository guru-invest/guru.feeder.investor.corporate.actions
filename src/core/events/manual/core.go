package manual

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/utils"
)

func BasicManualEvents(customers []mapper.Customer, corporateActions map[string][]mapper.CorporateAction) {
	var in_customers []string
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	chuncks := utils.ChunkSliceUtil{}.ChunkSlice(in_customers, 500)
	for _, item := range chuncks {
		ManualTransaction := repository.GetManualTransaction(item)
		ManualTransactionPersisterObject := []mapper.ManualTransaction{}
		for _, transaction := range ManualTransaction {

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
					ManualTransactionPersisterObject = append(ManualTransactionPersisterObject, ApplyBasicCorporateAction(transaction, corporate_action))
					continue
				}
			}
		}

		repository.UpdateManualTransaction(ManualTransactionPersisterObject)
	}

}
