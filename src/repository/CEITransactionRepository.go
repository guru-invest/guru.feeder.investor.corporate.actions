package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type CEITransactionRepository struct {
	_connection DatabaseConnection
}

func (h CEITransactionRepository) getCEITransactions() ([]mapper.CEITransaction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var cei_transaction []mapper.CEITransaction
	err := h._connection._databaseConnection.
		Select("id, customer_code, symbol, broker_id, quantity, price, amount, trade_date, post_event_quantity, post_event_price, post_event_symbol, event_factor, event_date, event_name").
		Order("trade_date asc").
		Find(&cei_transaction).
		Error

	if err != nil {
		return []mapper.CEITransaction{}, err
	}

	return cei_transaction, nil
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h CEITransactionRepository) updateCEITransactions(CEITransaction []mapper.CEITransaction) {
	h._connection.connect()
	defer h._connection.disconnect()

	for _, value := range CEITransaction {
		err := h._connection._databaseConnection.Save(&value).Error
		if err != nil {
			log.Println(err)
		}
	}
}

func GetCEITransaction() []mapper.CEITransaction {
	db := CEITransactionRepository{}
	cei_transaction, err := db.getCEITransactions()
	if err != nil {
		log.Println(err)
		return []mapper.CEITransaction{}
	}

	return cei_transaction
}

func UpdateCEITransaction(CEITransaction []mapper.CEITransaction) {
	db := CEITransactionRepository{}
	db.updateCEITransactions(CEITransaction)
}

var CEITransactionMap = map[string][]mapper.CEITransaction{}

func GetAllCEITransactions() map[string][]mapper.CEITransaction {

	if len(CEITransactionMap) == 0 {
		allCEITransactions := getAllCEITransactionsMap()
		for _, transaction := range allCEITransactions {
			mutex.Lock()
			CEITransactionMap[transaction.CustomerCode] = append(CEITransactionMap[transaction.CustomerCode], transaction)
			mutex.Unlock()
		}

		return CEITransactionMap
	}
	return CEITransactionMap
}

func getAllCEITransactionsMap() []mapper.CEITransaction {
	db := CEITransactionRepository{}
	cei_transaction, err := db.getCEITransactions()
	if err != nil {
		log.Println(err)
		return []mapper.CEITransaction{}
	}

	return cei_transaction
}
