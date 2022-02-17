package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type SymbolRepository struct {
	_connection DatabaseConnection
}

func (h SymbolRepository) getSymbols() ([]mapper.Symbol, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var symbol []mapper.Symbol
	err := h._connection._databaseConnection.Distinct("symbol").Find(&symbol, "event_date is null").Error
	if err != nil {
		return []mapper.Symbol{}, err
	}

	return symbol, nil
}

func GetSymbols() []mapper.Symbol {
	db := SymbolRepository{}
	symbols, err := db.getSymbols()
	if err != nil {
		log.Println(err)
		return []mapper.Symbol{}
	}

	return symbols
}
