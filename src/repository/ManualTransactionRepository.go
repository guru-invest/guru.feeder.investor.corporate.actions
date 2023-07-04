package repository

import (
	"fmt"
	"log"
	"sync"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

type ManualTransactionRepository struct {
	_connection DatabaseConnection
}

func (h ManualTransactionRepository) getManualTransactions(customers []string) ([]mapper.ManualTransaction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var manual_transaction []mapper.ManualTransaction

	result := h._connection._databaseConnection.
		Select("id, customer_code, symbol, broker_id, quantity, price, amount, side, trade_date, source_type, post_event_quantity, post_event_price, post_event_symbol, event_factor, event_date, event_name").
		Where("customer_code in ? and event_name <> ?", customers, constants.Bonus).
		Order("trade_date asc").
		Find(&manual_transaction)

	if result.Error != nil {
		fmt.Println(result.Error)
		return []mapper.ManualTransaction{}, result.Error
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

func (h ManualTransactionRepository) insertManualTransactions(manualTransaction []mapper.ManualTransaction, isStateLess bool) error {
	if isStateLess {
		h._connection.connectStateLess()
	} else {
		h._connection.connect()
	}

	defer h._connection.disconnect()

	var wg sync.WaitGroup
	for _, value := range manualTransaction {
		wg.Add(1)
		go func(valuer mapper.ManualTransaction) {
			defer wg.Done()
			err := h._connection._databaseConnection.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "hash_id"}},
				DoNothing: true,
			}).Create(&valuer).Error
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"Caller":        "guru.feeder.investor.corporate.actions.insertManualTransactions",
					"isStatless":    isStateLess,
					"CustomerCode:": valuer.CustomerCode,
					"Error":         err.Error(),
				}).Error("error insert manual proceeds")
			}
		}(value)

	}
	wg.Wait()
	return nil
}

func GetManualTransaction(customers []string) []mapper.ManualTransaction {
	db := ManualTransactionRepository{}
	manual_transaction, err := db.getManualTransactions(customers)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		return []mapper.ManualTransaction{}
	}

	return manual_transaction
}

func UpdateManualTransaction(manualTransaction []mapper.ManualTransaction) {
	db := ManualTransactionRepository{}
	db.updateManualTransactions(manualTransaction)
}

func InsertManualTransaction(manualTransaction []mapper.ManualTransaction, isStateLess bool) error {
	return ManualTransactionRepository{}.insertManualTransactions(manualTransaction, isStateLess)
}
