package interfaces

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
)

type ITransactionsRepository interface {
	GetTransactionsByCustomerCodes(customers []dtos.Customer) (*domain.TransactionList, error)
	UpdateTransactions(transactions domain.TransactionList) error
	GetSymbolsByCustomerCodes(customers []dtos.Customer) ([]dtos.Symbol, error)
}
