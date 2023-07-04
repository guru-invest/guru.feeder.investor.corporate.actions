package repository

import (
	"database/sql"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
)

type CorporateActionRepository struct {
	_connection DatabaseConnection
}

func (h CorporateActionRepository) getAllCorporateActions(asc_or_desc string) ([]mapper.CorporateAction, error) {
	h._connection.connect()
	defer h._connection.disconnect()
	var corporate_action []mapper.CorporateAction

	descriptions := []string{constants.Update, constants.Grouping, constants.Unfolding, constants.InterestOnEquity, constants.Dividend, constants.Income, constants.Bonus, constants.InterestOnEquityUp2Data, constants.Incorporation}

	err := h._connection._databaseConnection.
		Raw(`SELECT 
			ca.ticker, 
			ca.description, 
			ca.value, 
			ca.payment_date, 
			ca.com_date, 
			ca.target_ticker, 
			ca.calculated_factor, 
			ca.initial_date
		FROM 
			financial.corporate_actions ca
		WHERE
			ca.com_date > current_timestamp - interval '2 years' 
			and unaccent(ca.description) in @descriptions
		ORDER BY
			ca.com_date DESC`, sql.Named("descriptions", descriptions)).
		Find(&corporate_action).Error

	if err != nil {
		return []mapper.CorporateAction{}, err
	}

	return corporate_action, nil
}

func getAllCorporateActionsMap(asc_or_desc string) []mapper.CorporateAction {
	db := CorporateActionRepository{}
	corporate_actions, err := db.getAllCorporateActions(asc_or_desc)
	if err != nil {
		return []mapper.CorporateAction{}
	}

	return corporate_actions
}

//var mutex = &sync.Mutex{}

func GetCorporateActions(asc_or_desc string) map[string][]mapper.CorporateAction {
	corporateActionsMap := map[string][]mapper.CorporateAction{}
	allCorporateActions := getAllCorporateActionsMap(asc_or_desc)
	for _, value := range allCorporateActions {
		corporateActionsMap[value.Symbol] = append(corporateActionsMap[value.Symbol], value)
	}

	return corporateActionsMap
}
