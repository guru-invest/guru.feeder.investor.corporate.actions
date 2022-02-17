package repository

import (
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
)

type CorporateActionRepository struct {
	_connection DatabaseConnection
}

func (h CorporateActionRepository) GetCorporateActions(symbol string) ([]mapper.CorporateAction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var corporate_action []mapper.CorporateAction
	err := h._connection._databaseConnection.
		Select("description, com_date, target_ticker, calculated_factor").
		Order("com_date desc").
		Find(&corporate_action, "com_date > current_timestamp - interval '5 years' and description in (?, ?, ?) and ticker = ?",
			singleton.New().Update, singleton.New().Grouping, singleton.New().Unfolding, symbol).
		Error

	if err != nil {
		return []mapper.CorporateAction{}, err
	}

	return corporate_action, nil
}
