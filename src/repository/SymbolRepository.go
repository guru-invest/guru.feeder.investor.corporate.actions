package repository

import (
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type SymbolRepository struct {
	_connection DatabaseConnection
}

func (h SymbolRepository) GetSymbols() ([]mapper.Symbol, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var symbol []mapper.Symbol
	err := h._connection._databaseConnection.Distinct("symbol").Find(&symbol, "event_date is null").Error
	if err != nil {
		return []mapper.Symbol{}, err
	}

	return symbol, nil
}
