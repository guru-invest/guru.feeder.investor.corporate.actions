package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type CustomerRepository struct {
	_connection DatabaseConnection
}

func (h CustomerRepository) getOMSCustomers() ([]mapper.Customer, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var customer []mapper.Customer
	err := h._connection._databaseConnection.Table("wallet.oms_transactions").Distinct("customer_code").Find(&customer).Error
	if err != nil {
		return []mapper.Customer{}, err
	}

	return customer, nil
}

func GetOMSCustomers() []mapper.Customer {
	db := CustomerRepository{}
	customers, err := db.getOMSCustomers()
	if err != nil {
		log.Println(err)
		return []mapper.Customer{}
	}

	return customers
}

func (h CustomerRepository) getManualCustomers() ([]mapper.Customer, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var customer []mapper.Customer
	err := h._connection._databaseConnection.Table("wallet.manual_transactions").Distinct("customer_code").Find(&customer).Error
	if err != nil {
		return []mapper.Customer{}, err
	}

	return customer, nil
}

func GetManualCustomers() []mapper.Customer {
	db := CustomerRepository{}
	customers, err := db.getManualCustomers()
	if err != nil {
		log.Println(err)
		return []mapper.Customer{}
	}

	return customers
}

func (h CustomerRepository) getCEICustomers() ([]mapper.Customer, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var customer []mapper.Customer
	err := h._connection._databaseConnection.Table("wallet.cei_items_status").Distinct("customer_code").Where("execution_status = ?", "SUCCESS").Find(&customer).Error
	if err != nil {
		return []mapper.Customer{}, err
	}

	return customer, nil
}

func GetCEICustomers() []mapper.Customer {
	db := CustomerRepository{}
	customers, err := db.getCEICustomers()
	if err != nil {
		log.Println(err)
		return []mapper.Customer{}
	}

	return customers
}
