package repository

import (
	"log"
	"sync"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

type CorporateActionRepository struct {
	_connection DatabaseConnection
}

func (h CorporateActionRepository) getAllCorporateActions(asc_or_desc string) ([]mapper.CorporateAction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var corporate_action []mapper.CorporateAction
	err := h._connection._databaseConnection.
		Select("ticker, description, value, payment_date, com_date, target_ticker, calculated_factor").
		Order("com_date "+asc_or_desc).
		Find(&corporate_action, "com_date > current_timestamp - interval '5 years' and description in (?, ?, ?, ?, ?, ?, ?)",
			constants.Update, constants.Grouping, constants.Unfolding, constants.InterestOnEquity, constants.Dividend, constants.Income, constants.Bonus).
		Error

	if err != nil {
		return []mapper.CorporateAction{}, err
	}

	return corporate_action, nil
}

func getAllCorporateActionsMap(asc_or_desc string) []mapper.CorporateAction {
	db := CorporateActionRepository{}
	corporate_actions, err := db.getAllCorporateActions(asc_or_desc)
	if err != nil {
		log.Println(err)
		return []mapper.CorporateAction{}
	}

	return corporate_actions
}

var mutex = &sync.Mutex{}
var corporateActionsMap = map[string][]mapper.CorporateAction{}

func GetCorporateActions(asc_or_desc string) map[string][]mapper.CorporateAction {
	if len(corporateActionsMap) == 0 {
		allCorporateActions := getAllCorporateActionsMap(asc_or_desc)
		for _, value := range allCorporateActions {
			mutex.Lock()
			corporateActionsMap[value.Symbol] = append(corporateActionsMap[value.Symbol], value)
			mutex.Unlock()
		}

		return corporateActionsMap
	}
	return corporateActionsMap
}
