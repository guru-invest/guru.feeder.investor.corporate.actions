package repository

import (
	"log"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type OMSTransactionRepository struct {
	_connection DatabaseConnection
}

func (h OMSTransactionRepository) getOMSTransactions(symbol string, begin_date, end_date time.Time) ([]mapper.OMSTransaction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var oms_transaction []mapper.OMSTransaction
	err := h._connection._databaseConnection.
		Select("id, symbol, quantity, price, post_event_quantity, post_event_price, post_event_symbol, event_factor, event_date, event_name").
		Order("trade_date desc").
		Find(&oms_transaction, "symbol = ? and trade_date between ? and ?",
			symbol, begin_date, end_date).
		Error

	if err != nil {
		return []mapper.OMSTransaction{}, err
	}

	return oms_transaction, nil
}

func (h OMSTransactionRepository) updateOMSTransactions(OMSTransaction []mapper.OMSTransaction) {
	h._connection.connect()
	defer h._connection.disconnect()

	for _, value := range OMSTransaction {
		err := h._connection._databaseConnection.Save(&value).Error
		if err != nil {
			log.Println(err)
		}
	}
}

func GetOMSTransaction(symbol string, begin_date, end_date time.Time) []mapper.OMSTransaction {
	db := OMSTransactionRepository{}
	oms_transaction, err := db.getOMSTransactions(symbol, begin_date, end_date)
	if err != nil {
		log.Println(err)
		return []mapper.OMSTransaction{}
	}

	return oms_transaction
}

func UpdateOMSTransaction(OMSTransaction []mapper.OMSTransaction) {
	db := OMSTransactionRepository{}
	db.updateOMSTransactions(OMSTransaction)
}
