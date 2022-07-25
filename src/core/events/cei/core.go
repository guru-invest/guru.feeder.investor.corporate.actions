package cei

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/sirupsen/logrus"
)

func BasicCEIEvents(customers []mapper.Customer, corporateActions map[string][]mapper.CorporateAction, isStateLess bool) {
	CEITransaction := repository.GetCEITransaction(customers, isStateLess)
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

func ProceedsCEIEvents(corporateActions map[string][]mapper.CorporateAction, customers []mapper.Customer, symbols []mapper.Symbol, isStateLess bool) {
	CEITransactions := repository.GetAllCEITransactions(customers, isStateLess)
	CEIProceedPersisterObject := []mapper.CEIProceeds{}
	for _, customer := range customers {

		logrus.WithFields(logrus.Fields{
			"CustomerCode":    customer.CustomerCode,
			"CeiTransactions": CEITransactions,
		}).Info("preenche obj CEITransactions")

		for _, symbol := range symbols {
			CEIProceedPersisterObject = append(CEIProceedPersisterObject, ApplyProceedsCorporateAction(customer.CustomerCode, symbol.Name, CEITransactions, corporateActions)...)
			logrus.WithFields(logrus.Fields{
				"CustomerCode":              customer.CustomerCode,
				"Symbol":                    symbol.Name,
				"CEIProceedPersisterObject": CEIProceedPersisterObject,
			}).Info("For do ProceedsCEIEvents preenche obj CEIProceedPersisterObject")
		}
	}

	if len(CEIProceedPersisterObject) > 0 {
		//logrus.WithFields(logrus.Fields{}).Info("aqui tem que inserir na tabela")
		err := repository.InsertCEIProceeds(CEIProceedPersisterObject, isStateLess)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Caller":        "guru.feeder.investor.corporate.actions",
				"isStatless":    isStateLess,
				"CustomerCode:": customers[0].CustomerCode,
				"Error":         err.Error(),
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
