package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type CustomerRepository struct {
	_connection DatabaseConnection
}

func (h CustomerRepository) getCustomers() ([]mapper.Customer, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var customer []mapper.Customer
	err := h._connection._databaseConnection.Distinct("customer_code").Find(&customer).Error
	if err != nil {
		return []mapper.Customer{}, err
	}

	return customer, nil
}

func GetCustomers() []mapper.Customer {
	db := CustomerRepository{}
	customers, err := db.getCustomers()
	if err != nil {
		log.Println(err)
		return []mapper.Customer{}
	}

	return customers
}
