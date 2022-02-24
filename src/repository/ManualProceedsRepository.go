package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type ManualProceedsRepository struct {
	_connection DatabaseConnection
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h ManualProceedsRepository) insertManualProceeds(ManualProceeds []mapper.ManualProceeds) {
	h._connection.connect()
	defer h._connection.disconnect()

	err := h._connection._databaseConnection.Create(&ManualProceeds).Error
	if err != nil {
		log.Println(err)
	}

}

func InsertManualProceeds(ManualProceeds []mapper.ManualProceeds) {
	db := ManualProceedsRepository{}
	db.insertManualProceeds(ManualProceeds)
}
