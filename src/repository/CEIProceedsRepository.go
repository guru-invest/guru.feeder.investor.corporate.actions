package repository

import (
	"fmt"

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

	err := h._connection._databaseConnection.Clauses(clause.OnConflict{DoNothing: true}).Debug().CreateInBatches(&CEIProceeds, 300).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func InsertCEIProceeds(CEIProceeds []mapper.CEIProceeds, isStateLess bool) error {
	return CEIProceedsRepository{}.insertCEIProceeds(CEIProceeds, isStateLess)
}
