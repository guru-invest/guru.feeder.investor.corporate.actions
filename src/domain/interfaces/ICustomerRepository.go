package interfaces

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
)

type ICustomerRepository interface {
	GetAll() ([]dtos.Customer, error)
}
