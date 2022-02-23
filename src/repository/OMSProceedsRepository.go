package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type OMSProceedsRepository struct {
	_connection DatabaseConnection
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h OMSProceedsRepository) insertOMSProceeds(OMSProceeds []mapper.OMSProceeds) {
	h._connection.connect()
	defer h._connection.disconnect()

	err := h._connection._databaseConnection.Create(&OMSProceeds).Error
	if err != nil {
		log.Println(err)
	}

}

func InsertOMSProceeds(OMSProceeds []mapper.OMSProceeds) {
	db := OMSProceedsRepository{}
	db.insertOMSProceeds(OMSProceeds)
}
