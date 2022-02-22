package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type ManualTransactionRepository struct {
	_connection DatabaseConnection
}

func (h ManualTransactionRepository) getManualTransactions(symbol string) ([]mapper.ManualTransaction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var manual_transaction []mapper.ManualTransaction
	err := h._connection._databaseConnection.
		Select("id, symbol, quantity, price, trade_date, post_event_quantity, post_event_price, post_event_symbol, event_factor, event_date, event_name").
		Order("trade_date asc").
		Find(&manual_transaction, "symbol = ?",
			symbol).
		Error

	if err != nil {
		return []mapper.ManualTransaction{}, err
	}

	return manual_transaction, nil
}

func (h ManualTransactionRepository) updateManualTransactions(ManualTransaction []mapper.ManualTransaction) {
	h._connection.connect()
	defer h._connection.disconnect()

	for _, value := range ManualTransaction {
		err := h._connection._databaseConnection.Save(&value).Error
		if err != nil {
			log.Println(err)
		}
	}
}

func GetManualTransaction(symbol string) []mapper.ManualTransaction {
	db := ManualTransactionRepository{}
	manual_transaction, err := db.getManualTransactions(symbol)
	if err != nil {
		log.Println(err)
		return []mapper.ManualTransaction{}
	}

	return manual_transaction
}

func UpdateManualTransaction(ManualTransaction []mapper.ManualTransaction) {
	db := ManualTransactionRepository{}
	db.updateManualTransactions(ManualTransaction)
}