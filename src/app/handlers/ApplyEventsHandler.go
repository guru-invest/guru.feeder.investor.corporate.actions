package handlers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/interfaces"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/repository"
)

type ApplyEventsHandler struct {
	_corporateActionsRepository interfaces.ICorporateActionRepository
	_customerRepository         interfaces.ICustomerRepository
	_transactionRepository      interfaces.ITransactionsRepository
}

var CorporateActionsAsc map[string][]mapper.CorporateAction
var CorporateActionsDesc map[string][]mapper.CorporateAction
var hashLogID string

func (t ApplyEventsHandler) NewApplyEventsHandler(customerCode string) error {
	return ApplyEventsHandler{
		_corporateActionsRepository: repository.NewCorporateActionRepository(),
		_customerRepository:         repository.NewCustomerRepository(),
		_transactionRepository:      repository.NewTransactionsRepository(),
	}.handler(customerCode)
}

func (t ApplyEventsHandler) handler(customerCode string) error {
	start := time.Now()

	// getCorporateActions, err := t._corporateActionsRepository.GetAll()
	// if err != nil {
	// 	return err
	// }

	t.applyEvents(customerCode)
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
	return nil
}

func (t ApplyEventsHandler) applyEvents(customerCode string) {
	id := uuid.New()
	hashLogID = id.String()

	corporateActions, err := t._corporateActionsRepository.GetAll()
	if err != nil {

	}

	OMSCustomers := []dtos.Customer{}
	if customerCode == constants.AllCustomers {
		OMSCustomers, err = t._customerRepository.GetAll()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"HashID": hashLogID,
				"Error":  err,
			}).Error("erro consultando GetOMSCustomers")
		}
	} else {
		OMSCustomers = []dtos.Customer{
			{
				CustomerCode: customerCode,
				CreatedAT:    time.Now().String(),
			},
		}
	}
	t.basicEvents(OMSCustomers, corporateActions)
	//err := applyAllEventsOMS(customerCode)
	//if err != nil {
	//	return
	//}

	walletConnector := repository.NewWalletConnector()
	walletConnector.ResyncAVGOMS()

}

func applyAllEventsOMS(customerCode string) error {
	//var wg sync.WaitGroup

	// logrus.WithFields(logrus.Fields{
	// 	"HashID": hashLogID,
	// }).Info("Inicia Eventos OMS")

	// OMSCustomers := []mapper.Customer{}
	// if customerCode == constants.AllCustomers {
	// 	GetOMSCustomers, err := repository.GetOMSCustomers()
	// 	if err != nil {
	// 		logrus.WithFields(logrus.Fields{
	// 			"HashID": hashLogID,
	// 			"Error":  err,
	// 		}).Error("erro consultando GetOMSCustomers")
	// 		return err
	// 	}
	// 	if OMSCustomers == nil {
	// 		logrus.WithFields(logrus.Fields{
	// 			"HashID": hashLogID,
	// 		}).Info("GetOMSCustomers not found")
	// 		return nil
	// 	}
	// 	OMSCustomers = GetOMSCustomers
	// } else {
	// 	Customer := mapper.Customer{
	// 		CustomerCode: customerCode,
	// 		CreatedAT:    time.Now().String(),
	// 	}

	// 	OMSCustomers = []mapper.Customer{Customer}
	// }

	// // wg.Add(2)
	// // go func() {
	// // 	defer wg.Done()
	// oms.BasicOMSEvents(OMSCustomers, CorporateActionsDesc)
	// //}()

	// // go func() {
	// // 	defer wg.Done()
	// // 	doProceedsOMSEvents(false, OMSCustomers)
	// // }()

	// // wg.Wait()

	logrus.WithFields(logrus.Fields{
		"HashID": hashLogID,
	}).Info("Finaliza Eventos OMS")

	return nil
}

func (t ApplyEventsHandler) basicEvents(customers []dtos.Customer, corporateActions *domain.CorporateActionList) {
	transactions, err := t._transactionRepository.GetTransactionsByCustomerCodes(customers)
	if err != nil {
		return
	}
	transactionPersisterObject := domain.TransactionList{}
	for _, transaction := range transactions.List {

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
				transactionPersisterObject.List = append(transactionPersisterObject.List, t.applyBasicCorporateAction(transaction, corporate_action))
				continue
			}
		}
	}

	t._transactionRepository.UpdateTransactions(transactionPersisterObject)
}

func (t ApplyEventsHandler) applyBasicCorporateAction(transaction domain.Transaction, corporate_action domain.CorporateAction) domain.Transaction {
	targetTicket := corporate_action.TargetTicker
	if corporate_action.TargetTicker == "" {
		targetTicket = corporate_action.Symbol
	}
	transaction.EventFactor = corporate_action.CalculatedFactor
	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if corporate_action.Description == constants.Update {
		transaction.EventFactor = decimal.NewFromInt32(1)
	}

	transaction.EventName = corporate_action.Description
	transaction.PostEventSymbol = targetTicket
	transaction.EventDate = corporate_action.ComDate
	transaction.PostEventQuantity = transaction.PostEventQuantity.Div(transaction.EventFactor)
	transaction.PostEventPrice = transaction.Amount.Div(transaction.PostEventQuantity).Truncate(2)

	// Processo cumulativo
	// OMSTransaction.Quantity = int(utils.Truncate(OMSTransaction.PostEventQuantity, 0))
	// OMSTransaction.Price = OMSTransaction.PostEventPrice

	return transaction
}
