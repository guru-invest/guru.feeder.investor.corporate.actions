package repository

import (
	"log"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
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
	err := h._connection._databaseConnection.
		Table("wallet.investor_sync_historical").
		Distinct("customer_code").
		Where("execution_status = ? AND event_type = ?", "SUCCESS", "movementEquities").
		Order("created_at DESC").
		Find(&customer).Error
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
