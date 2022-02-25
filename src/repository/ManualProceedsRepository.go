package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"gorm.io/gorm/clause"
)

type ManualProceedsRepository struct {
	_connection DatabaseConnection
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h ManualProceedsRepository) insertManualProceeds(manualProceeds []mapper.ManualProceeds) {
	h._connection.connect()
	defer h._connection.disconnect()

	err := h._connection._databaseConnection.Clauses(clause.OnConflict{DoNothing: true}).Create(&manualProceeds).Error
	if err != nil {
		log.Println(err)
	}

}

func InsertManualProceeds(manualProceeds []mapper.ManualProceeds) {
	db := ManualProceedsRepository{}
	db.insertManualProceeds(manualProceeds)
}
