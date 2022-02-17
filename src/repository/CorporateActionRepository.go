package repository

import (
	"log"
	"sync"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
)

type CorporateActionRepository struct {
	_connection DatabaseConnection
}

func (h CorporateActionRepository) getAllCorporateActions() ([]mapper.CorporateAction, error) {
	h._connection.connect()
	defer h._connection.disconnect()

	var corporate_action []mapper.CorporateAction
	err := h._connection._databaseConnection.
		Select("ticker, description, com_date, target_ticker, calculated_factor").
		Order("com_date desc").
		Find(&corporate_action, "com_date > current_timestamp - interval '5 years' and description in (?, ?, ?)",
			singleton.New().Update, singleton.New().Grouping, singleton.New().Unfolding).
		Error

	if err != nil {
		return []mapper.CorporateAction{}, err
	}

	return corporate_action, nil
}

func getAllCorporateActionsMap() []mapper.CorporateAction {
	db := CorporateActionRepository{}
	corporate_actions, err := db.getAllCorporateActions()
	if err != nil {
		log.Println(err)
		return []mapper.CorporateAction{}
	}

	return corporate_actions
}

var mutex = &sync.Mutex{}
var corporateActionsMap = map[string][]mapper.CorporateAction{}

func GetCorporateActions(symbol string) []mapper.CorporateAction {

	if len(corporateActionsMap) == 0 {
		allCorporateActions := getAllCorporateActionsMap()
		for _, value := range allCorporateActions {
			mutex.Lock()
			corporateActionsMap[value.Symbol] = append(corporateActionsMap[value.Symbol], value)
			mutex.Unlock()
		}
		// TODO - Porque quando o map não existe, ele não passa no return de fora ?
		return corporateActionsMap[symbol]
	}

	return corporateActionsMap[symbol]
}
