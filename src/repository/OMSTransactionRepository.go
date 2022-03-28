package repository

import (
	"log"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
)

type OMSTransactionRepository struct {
	_connection DatabaseConnection
}

func (h OMSTransactionRepository) getOMSTransactions(customers []mapper.Customer) ([]mapper.OMSTransaction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var oms_transaction []mapper.OMSTransaction
	var in_customers []string
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	err := h._connection._databaseConnection.
		Select("id, customer_code, broker_id, symbol, quantity, price, amount, side, trade_date, post_event_quantity, post_event_price, post_event_symbol, event_factor, event_date, event_name").
		Where("customer_code in ?", in_customers).
		Order("trade_date asc").
		Find(&oms_transaction).
		Error

	if err != nil {
		return []mapper.OMSTransaction{}, err
	}

	return oms_transaction, nil
}

// TODO - Não deveria estar persistindo dados aqui no repository
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

func GetOMSTransaction(customers []mapper.Customer) []mapper.OMSTransaction {
	db := OMSTransactionRepository{}
	oms_transaction, err := db.getOMSTransactions(customers)
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

var OMSTransactionMap = map[string][]mapper.OMSTransaction{}

func GetAllOMSTransactions(customers []mapper.Customer) map[string][]mapper.OMSTransaction {

	if len(OMSTransactionMap) == 0 {
		allOMSTransactions := getAllOMSTransactionsMap(customers)
		for _, transaction := range allOMSTransactions {
			mutex.Lock()
			OMSTransactionMap[transaction.CustomerCode] = append(OMSTransactionMap[transaction.CustomerCode], transaction)
			mutex.Unlock()
		}

		return OMSTransactionMap
	}
	return OMSTransactionMap
}

func getAllOMSTransactionsMap(customers []mapper.Customer) []mapper.OMSTransaction {
	db := OMSTransactionRepository{}
	oms_transaction, err := db.getOMSTransactions(customers)
	if err != nil {
		log.Println(err)
		return []mapper.OMSTransaction{}
	}

	return oms_transaction
}
