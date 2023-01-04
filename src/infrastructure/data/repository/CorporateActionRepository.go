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

func (t CorporateActionRepository) GetAll() (*domain.CorporateActionList, error) {
	getCorporateAction := []mapper.CorporateAction{}
	descriptions := []string{constants.Update, constants.Grouping, constants.Unfolding, constants.InterestOnEquity, constants.Dividend, constants.Income, constants.Bonus}

	err := t._connection.DatabaseConnection.
		Raw(statements.CorporateActionsGetAll, sql.Named("descriptions", descriptions)).
		Find(&getCorporateAction).Error
	if err != nil {
		return nil, err
	}

	return mapper.CorporateAction{}.ToDomain(getCorporateAction), nil
}
