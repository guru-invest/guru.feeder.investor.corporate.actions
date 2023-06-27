package repository

import (
	"sync"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
)

type CorporateActionRepository struct {
	_connection DatabaseConnection
}

func (h CorporateActionRepository) getAllCorporateActions(asc_or_desc string) (*[]domain.CorporateActions, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	corporateActionList := mapper.CorporateActionList{}
	err := h._connection._databaseConnection.
		Select("ticker, description, value, payment_date, com_date, target_ticker, calculated_factor, initial_date").
		Order("com_date "+asc_or_desc).
		Find(&corporateActionList.CorporateActionList, "com_date > current_timestamp - interval '5 years' and description in (?, ?, ?, ?, ?, ?, ?)",
			constants.Update, constants.Grouping, constants.Unfolding, constants.InterestOnEquity, constants.Dividend, constants.Income, constants.Bonus).
		Error

	if err != nil {
		return nil, err
	}

	return corporateActionList.ToDomain(), nil
}

// func getAllCorporateActionsMap(asc_or_desc string) []mapper.CorporateAction {
// 	db := CorporateActionRepository{}
// 	corporate_actions, err := db.getAllCorporateActions(asc_or_desc)
// 	if err != nil {
// 		log.Println(err)
// 		return []mapper.CorporateAction{}
// 	}

// 	return corporate_actions
// }

var mutex = &sync.Mutex{}

// var corporateActionsMap = map[string][]mapper.CorporateAction{}

func GetCorporateActions(asc_or_desc string) (*[]domain.CorporateActions, error) {
	return CorporateActionRepository{}.getAllCorporateActions(asc_or_desc)
}
