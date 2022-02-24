package repository

import (
	"log"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type CEIProceedsRepository struct {
	_connection DatabaseConnection
}

// TODO - Não deveria estar persistindo dados aqui no repository
func (h CEIProceedsRepository) insertCEIProceeds(CEIProceeds []mapper.CEIProceeds) {
	h._connection.connect()
	defer h._connection.disconnect()

	err := h._connection._databaseConnection.Create(&CEIProceeds).Error
	if err != nil {
		log.Println(err)
	}

}

func InsertCEIProceeds(CEIProceeds []mapper.CEIProceeds) {
	db := CEIProceedsRepository{}
	db.insertCEIProceeds(CEIProceeds)
}