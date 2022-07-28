package repository

import (
	"fmt"

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

func GetOMSCustomers() ([]mapper.Customer, error) {
	db := CustomerRepository{}
	customers, err := db.getOMSCustomers()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (h CustomerRepository) getManualCustomers() ([]mapper.Customer, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var customer []mapper.Customer
	err := h._connection._databaseConnection.Table("wallet.manual_transactions").Distinct("customer_code").Find(&customer).Error
	if err != nil {
		fmt.Println(err)
		return []mapper.Customer{}, err
	}

	return customer, nil
}

func GetManualCustomers() ([]mapper.Customer, error) {
	db := CustomerRepository{}
	customers, err := db.getManualCustomers()
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (h CustomerRepository) getCEICustomers() ([]mapper.Customer, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var customer []mapper.Customer
	err := h._connection._databaseConnection.
		Table("wallet.investor_sync_historical").
		Distinct("customer_code, MAX(created_at) as created_at").
		Where("execution_status = ? AND event_type = ?", "SUCCESS", "movementEquities").
		Group("customer_code").
		Order("MAX(created_at) DESC").
		Scan(&customer).Error
	if err != nil {
		return []mapper.Customer{}, err
	}

	return customer, nil
}

func GetCEICustomers() ([]mapper.Customer, error) {
	db := CustomerRepository{}
	customers, err := db.getCEICustomers()
	if err != nil {
		return nil, err
	}

	return customers, nil
}
