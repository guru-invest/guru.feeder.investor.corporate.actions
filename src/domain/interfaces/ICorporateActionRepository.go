package interfaces

import "github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"

type ICorporateActionRepository interface {
	GetAll() (*domain.CorporateActionList, error)
}
