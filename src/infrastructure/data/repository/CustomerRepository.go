package repository

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/dtos"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/interfaces"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/providers"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/statements"
)

type CustomerRepository struct {
	_connection providers.DatabaseConnection
}

func NewCustomerRepository() interfaces.ICustomerRepository {
	connection := CustomerRepository{
		_connection: providers.DatabaseConnection{},
	}
	connection._connection.Connect()
	return &connection
}

func (t CustomerRepository) GetAll() ([]dtos.Customer, error) {
	customers := []dtos.Customer{}

	err := t._connection.DatabaseConnection.
		Raw(statements.CustomerGetAll).
		Find(&customers).Error()
	if err != nil {
		return nil, err
	}

	return customers, nil
}
