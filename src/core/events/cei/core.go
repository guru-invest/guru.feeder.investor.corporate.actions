package cei

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/utils"
	"github.com/sirupsen/logrus"
)

func BasicCEIEvents(customers []mapper.Customer, corporateActions map[string][]mapper.CorporateAction, isStateLess bool) {

	var in_customers []string
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	chuncks := utils.ChunkSliceUtil{}.ChunkSlice(in_customers, 200)
	for _, item := range chuncks {
		CEITransaction := repository.GetCEITransaction(item, isStateLess)
		CEITransactionPersisterObject := []mapper.CEITransaction{}

		for _, transaction := range CEITransaction {

			for _, corporate_action := range corporateActions[transaction.Symbol] {

				// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
				if transaction.TradeDate.After(corporate_action.InitialDate) && !transaction.TradeDate.Equal(corporate_action.InitialDate) {
					continue
				}

				// Se a data de EventDate for maior, significa que eventos corporativos com datas inferiores, já foram aplicados (Eventos com ano 2001 sao eventos com data default e devem ser consideraros pois é a primeira vez)
				if corporate_action.InitialDate.After(transaction.EventDate) && transaction.EventDate.Year() > 2001 {
					continue
				}

				// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
				if corporate_action.IsBasic() {
					CEITransactionPersisterObject = append(CEITransactionPersisterObject, ApplyBasicCorporateAction(transaction, corporate_action))
					continue
				}
			}
		}

		repository.UpdateCEITransaction(CEITransactionPersisterObject, isStateLess)
	}

}

func ProceedsCEIEvents(corporateActions map[string][]mapper.CorporateAction, customers []mapper.Customer, symbols []mapper.Symbol, isStateLess bool) {

	var in_customers []string 
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	chuncks := utils.ChunkSliceUtil{}.ChunkSlice(in_customers, 200)
	for _, item := range chuncks {
		CEITransactions := repository.GetAllCEITransactions(item, isStateLess)
		CEIProceedPersisterObject := []mapper.CEIProceeds{}
		for _, customer := range customers {
			for _, symbol := range symbols {
				CEIProceedPersisterObject = append(CEIProceedPersisterObject, ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, CEITransactions, corporateActions)...)
			}
		}

		if len(CEIProceedPersisterObject) > 0 {
			err := repository.InsertCEIProceeds(CEIProceedPersisterObject, isStateLess)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"Caller":     "guru.feeder.investor.corporate.actions",
					"isStatless": isStateLess,
					"Error":      err.Error(),
				}).Error("error insert investor proceeds")
			}
		}

		ManualTransactions := []mapper.ManualTransaction{}
		for _, proceed := range CEIProceedPersisterObject {
			if proceed.Event == constants.Bonus {
				ManualTransaction := mapper.ManualTransaction{}
				ManualTransactions = append(ManualTransactions, manual.ApplyInheritedBonusActionCEI(ManualTransaction, proceed))
			}
		}

		if len(ManualTransactions) > 0 {
			repository.InsertManualTransaction(ManualTransactions, isStateLess)
		}
	}

}
