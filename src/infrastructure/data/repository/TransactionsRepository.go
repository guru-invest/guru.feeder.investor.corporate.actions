package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/interfaces"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/providers"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/statements"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/repository/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/utils"
)

type TransactionsRepository struct {
	_connection providers.DatabaseConnection
}

func NewTransactionsRepository() interfaces.ITransactionsRepository {
	connection := TransactionsRepository{
		_connection: providers.DatabaseConnection{},
	}
	connection._connection.Connect()
	return &connection
}

func (t TransactionsRepository) GetTransactionsByCustomerCodes(customers []dtos.Customer) (*domain.TransactionList, error) {

	var oms_transaction []mapper.Transaction
	var in_customers []string
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	chuncks := utils.ChunkSliceUtil{}.ChunkSlice(in_customers, 1000)

	for _, customer := range chuncks {
		var internal_oms_transaction []mapper.Transaction

		err := t._connection.DatabaseConnection.
			Raw(statements.GetTransactionsByCustomerCodes, sql.Named("customerCodes", customer)).
			Find(&internal_oms_transaction).
			Error()

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		oms_transaction = append(oms_transaction, internal_oms_transaction...)
	}

	return mapper.Transaction{}.ToDomain(oms_transaction), nil
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (t TransactionsRepository) updateOMSTransactions(OMSTransaction []mapper.OMSTransaction) {
	for _, value := range OMSTransaction {
		err := t._connection.DatabaseConnection.Save(&value).Error
		if err != nil {
			fmt.Println(err)
			log.Println(err)
		}
	}
}

func (t TransactionsRepository) UpdateOMSTransaction(OMSTransaction domain.TransactionList) {
	db := TransactionsRepository{}
	db.updateOMSTransactions(OMSTransaction)
}

func GetAllOMSTransactions(customers []mapper.Customer) map[string][]mapper.OMSTransaction {
	var OMSTransactionMap = map[string][]mapper.OMSTransaction{}
	if len(OMSTransactionMap) == 0 {
		allOMSTransactions := getAllOMSTransactionsMap(customers)
		for _, transaction := range allOMSTransactions {
			mutex.Lock()
			OMSTransactionMap[transaction.CustomerCode] = append(OMSTransactionMap[transaction.CustomerCode], transaction)
			mutex.Unlock()
		}

		return OMSTransactionMap
	}
	return OMSTransactionMap
}

func getAllOMSTransactionsMap(customers []mapper.Customer) []mapper.OMSTransaction {
	db := TransactionsRepository{}
	oms_transaction, err := db.getOMSTransactions(customers)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return []mapper.OMSTransaction{}
	}

	return oms_transaction
}
