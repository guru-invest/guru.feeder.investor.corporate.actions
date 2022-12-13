package interfaces

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
)

type ITransactionsRepository interface {
	GetTransactionsByCustomerCodes(customers []dtos.Customer) ([]dtos.Transaction, error)
}
