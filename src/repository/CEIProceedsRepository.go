package repository

import (
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"gorm.io/gorm/clause"
)

type CEIProceedsRepository struct {
	_connection DatabaseConnection
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h CEIProceedsRepository) insertCEIProceeds(CEIProceeds []mapper.CEIProceeds, isStateLess bool) error {
	if isStateLess {
		h._connection.connectStateLess()
	} else {
		h._connection.connect()
	}
	defer h._connection.disconnect()

	err := h._connection._databaseConnection.Clauses(clause.OnConflict{DoNothing: true}).Create(&CEIProceeds).Error
	if err != nil {
		return err
	}
	return nil
}

func InsertCEIProceeds(CEIProceeds []mapper.CEIProceeds, isStateLess bool) error {
	return CEIProceedsRepository{}.insertCEIProceeds(CEIProceeds, isStateLess)
}
