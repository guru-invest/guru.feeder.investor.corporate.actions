package repository

import (
	"database/sql"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain/interfaces"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/providers"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/infrastructure/data/statements"
)

type CorporateActionRepository struct {
	_connection providers.DatabaseConnection
}

func NewCorporateActionRepository() interfaces.ICorporateActionRepository {
	connection := CorporateActionRepository{
		_connection: providers.DatabaseConnection{},
	}
	connection._connection.Connect()
	return &connection
}

// func (h CorporateActionRepository) getAllCorporateActions(asc_or_desc string) ([]mapper.CorporateAction, error) {

// 	var corporate_action []mapper.CorporateAction
// 	err := h._connection._databaseConnection.
// 		Select("ticker, description, value, payment_date, com_date, target_ticker, calculated_factor, initial_date").
// 		Order("com_date "+asc_or_desc).
// 		Find(&corporate_action, "com_date > current_timestamp - interval '5 years' and description in (?, ?, ?, ?, ?, ?, ?)",
// 			constants.Update, constants.Grouping, constants.Unfolding, constants.InterestOnEquity, constants.Dividend, constants.Income, constants.Bonus).
// 		Error

// 	if err != nil {
// 		return []mapper.CorporateAction{}, err
// 	}

// 	return corporate_action, nil
// }

// func getAllCorporateActionsMap(asc_or_desc string) []mapper.CorporateAction {
// 	db := CorporateActionRepository{}
// 	corporate_actions, err := db.getAllCorporateActions(asc_or_desc)
// 	if err != nil {
// 		log.Println(err)
// 		return []mapper.CorporateAction{}
// 	}

// 	return corporate_actions
// }

// func (h CorporateActionRepository) GetCorporateActions(asc_or_desc string) (map[string]domain.CorporateActionList, error) {
// 	var corporateActionsMap = map[string]domain.CorporateActionList{}

// 	allCorporateActions, err := h.getAllCorporateActions(asc_or_desc)
// 	if err != nil {
// 		return nil, err
// 	}
// 	for _, value := range allCorporateActions {
// 		mutex.Lock()
// 		corporateActionsMap[value.Symbol] = append(corporateActionsMap[value.Symbol], value)
// 		mutex.Unlock()
// 	}

// 	return corporateActionsMap, nil
// }

func (t CorporateActionRepository) GetAll() (*domain.CorporateActionList, error) {
	getCorporateAction := []mapper.CorporateAction{}
	descriptions := []string{constants.Update, constants.Grouping, constants.Unfolding, constants.InterestOnEquity, constants.Dividend, constants.Income, constants.Bonus}

	err := t._connection.DatabaseConnection.
		Raw(statements.CorporateActionsGetAll, sql.Named("descriptions", descriptions)).
		Find(&getCorporateAction).Error()
	if err != nil {
		return nil, err
	}

	return mapper.CorporateAction{}.ToDomain(getCorporateAction), nil
}
