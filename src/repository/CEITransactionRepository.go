package repository

import (
	"fmt"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
)

type CEITransactionRepository struct {
	_connection DatabaseConnection
}

func (h CEITransactionRepository) getCEITransactions(customers []string, isStateLess bool) ([]mapper.CEITransaction, error) {
	if isStateLess {
		h._connection.connectStateLess()
	} else {
		h._connection.connect()
	}
	defer h._connection.disconnect()

	var cei_transaction []mapper.CEITransaction

	err := h._connection._databaseConnection.
		Select("id, customer_code, symbol, broker_id, quantity, price, amount, side, trade_date, post_event_quantity, post_event_price, post_event_symbol, event_factor, event_date, event_name").
		Where("customer_code in ?", customers).
		Where("movement_type = ?", "Assets-Trading").
		Order("trade_date asc").
		Find(&cei_transaction).Error

	if err != nil {
		fmt.Println(err)
		return []mapper.CEITransaction{}, err
	}

	return cei_transaction, nil
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h CEITransactionRepository) updateCEITransactions(CEITransaction []mapper.CEITransaction, isStateLess bool) {

	if isStateLess {
		h._connection.connectStateLess()
	} else {
		h._connection.connect()
	}

	defer h._connection.disconnect()

	for _, value := range CEITransaction {
		err := h._connection._databaseConnection.Save(&value).Error
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetCEITransaction(customers []string, isStateLess bool) []mapper.CEITransaction {
	db := CEITransactionRepository{}
	cei_transaction, err := db.getCEITransactions(customers, isStateLess)
	if err != nil {
		fmt.Println(err)
		return []mapper.CEITransaction{}
	}

	return cei_transaction
}

func UpdateCEITransaction(CEITransaction []mapper.CEITransaction, isStateLess bool) {
	db := CEITransactionRepository{}
	db.updateCEITransactions(CEITransaction, isStateLess)
}

func GetAllCEITransactions(customers []string, isStateLess bool) map[string][]mapper.CEITransaction {
	var CEITransactionMap = map[string][]mapper.CEITransaction{}

	if len(CEITransactionMap) == 0 {
		allCEITransactions := getAllCEITransactionsMap(customers, isStateLess)
		for _, transaction := range allCEITransactions {
			mutex.Lock()
			CEITransactionMap[transaction.CustomerCode] = append(CEITransactionMap[transaction.CustomerCode], transaction)
			mutex.Unlock()
		}

		return CEITransactionMap
	}
	return CEITransactionMap
}

func getAllCEITransactionsMap(customers []string, isStateLess bool) []mapper.CEITransaction {
	db := CEITransactionRepository{}
	cei_transaction, err := db.getCEITransactions(customers, isStateLess)
	if err != nil {
		fmt.Println(err)
		return []mapper.CEITransaction{}
	}

	return cei_transaction
}
