package repository

import (
	"log"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"gorm.io/gorm/clause"
)

type CEIProceedsRepository struct {
	_connection DatabaseConnection
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h CEIProceedsRepository) insertCEIProceeds(CEIProceeds []mapper.CEIProceeds, isStateLess bool) {
	if isStateLess {
		h._connection.connectStateLess()
	} else {
		h._connection.connect()
	}
	defer h._connection.disconnect()

	err := h._connection._databaseConnection.Clauses(clause.OnConflict{DoNothing: true}).Create(&CEIProceeds).Error
	if err != nil {
		log.Println(err)
	}

}

func InsertCEIProceeds(CEIProceeds []mapper.CEIProceeds, isStateLess bool) {
	db := CEIProceedsRepository{}
	db.insertCEIProceeds(CEIProceeds, isStateLess)
}
