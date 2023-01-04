package repository

import (
	"database/sql"
	"fmt"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/utils"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/interfaces"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/providers"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/statements"
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
			Error

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		oms_transaction = append(oms_transaction, internal_oms_transaction...)
	}

	return mapper.Transaction{}.ToDomain(oms_transaction), nil
}

func (t TransactionsRepository) UpdateTransactions(transactions domain.TransactionList) error {

	for _, value := range transactions.List {
		err := t._connection.DatabaseConnection.Save(&value).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (t TransactionsRepository) GetSymbolsByCustomerCodes(customers []dtos.Customer) ([]dtos.Symbol, error) {
	var symbol []dtos.Symbol

	var in_customers []string
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	err := t._connection.DatabaseConnection.Table("wallet.oms_transactions").Distinct("symbol").Where("customer_code in ?", in_customers).Find(&symbol).Error
	if err != nil {
		return []dtos.Symbol{}, err
	}

	return symbol, nil
}
