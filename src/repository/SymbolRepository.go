package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type SymbolRepository struct {
	_connection DatabaseConnection
}

func (h SymbolRepository) getSymbols(customers []mapper.Customer) ([]mapper.Symbol, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var symbol []mapper.Symbol

	var in_customers []string
	for _, value := range customers {
		in_customers = append(in_customers, value.CustomerCode)
	}

	err := h._connection._databaseConnection.Distinct("symbol").Where("customer_code in ?", in_customers).Find(&symbol).Error
	if err != nil {
		return []mapper.Symbol{}, err
	}

	return symbol, nil
}

func GetSymbols(customers []mapper.Customer) []mapper.Symbol {
	db := SymbolRepository{}
	symbols, err := db.getSymbols(customers)
	if err != nil {
		log.Println(err)
		return []mapper.Symbol{}
	}

	return symbols
}
