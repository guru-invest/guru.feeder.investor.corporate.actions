package repository

import (
	"log"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type OMSTransactionRepository struct {
	_connection DatabaseConnection
}

func (h OMSTransactionRepository) getOMSTransactions(symbol, target_symbol, event string, begin_date, end_date time.Time) ([]mapper.OMSTransaction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var oms_transaction []mapper.OMSTransaction
	err := h._connection._databaseConnection.
		Select("id, symbol, quantity, price, post_event_quantity, post_event_price, ? as post_event_symbol, event_factor, event_date, ? as event_name", target_symbol, event).
		Order("trade_date desc").
		Find(&oms_transaction, "symbol = ? and trade_date between ? and ?",
			symbol, begin_date, end_date).
		Error

	if err != nil {
		return []mapper.OMSTransaction{}, err
	}

	return oms_transaction, nil
}

func GetOMSTransaction(symbol, target_symbol, event string, begin_date, end_date time.Time) []mapper.OMSTransaction {
	db := OMSTransactionRepository{}
	oms_transaction, err := db.getOMSTransactions(symbol, target_symbol, event, begin_date, end_date)
	if err != nil {
		log.Println(err)
		return []mapper.OMSTransaction{}
	}

	return oms_transaction
}
