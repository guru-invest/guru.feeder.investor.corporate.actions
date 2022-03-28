package repository

import (
	"log"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"gorm.io/gorm/clause"
)

type ManualTransactionRepository struct {
	_connection DatabaseConnection
}

func (h ManualTransactionRepository) getManualTransactions(customers []mapper.Customer) ([]mapper.ManualTransaction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var manual_transaction []mapper.ManualTransaction
	var in_customers []string
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	err := h._connection._databaseConnection.
		Select("id, customer_code, symbol, broker_id, quantity, price, amount, side, trade_date, source_type, post_event_quantity, post_event_price, post_event_symbol, event_factor, event_date, event_name").
		Where("customer_code in ? and event_name <> ?", in_customers, constants.Bonus).
		Order("trade_date asc").
		Find(&manual_transaction).
		Error

	if err != nil {
		return []mapper.ManualTransaction{}, err
	}

	return manual_transaction, nil
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h ManualTransactionRepository) updateManualTransactions(manualTransaction []mapper.ManualTransaction) {
	h._connection.connect()
	defer h._connection.disconnect()

	for _, value := range manualTransaction {
		err := h._connection._databaseConnection.Save(&value).Error
		if err != nil {
			log.Println(err)
		}
	}
}

func (h ManualTransactionRepository) insertManualTransactions(manualTransaction []mapper.ManualTransaction) {
	h._connection.connect()
	defer h._connection.disconnect()

	for _, value := range manualTransaction {
		err := h._connection._databaseConnection.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "hash_id"}},
			DoNothing: true,
		}).Create(&value).Error
		if err != nil {
			log.Println(err)
		}
	}
}

func GetManualTransaction(customers []mapper.Customer) []mapper.ManualTransaction {
	db := ManualTransactionRepository{}
	manual_transaction, err := db.getManualTransactions(customers)
	if err != nil {
		log.Println(err)
		return []mapper.ManualTransaction{}
	}

	return manual_transaction
}

func UpdateManualTransaction(manualTransaction []mapper.ManualTransaction) {
	db := ManualTransactionRepository{}
	db.updateManualTransactions(manualTransaction)
}

func InsertManualTransaction(manualTransaction []mapper.ManualTransaction) {
	db := ManualTransactionRepository{}
	db.insertManualTransactions(manualTransaction)
}
