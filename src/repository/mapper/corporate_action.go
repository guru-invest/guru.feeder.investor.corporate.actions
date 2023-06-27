package mapper

import (
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/domain"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/utils"
)

func (h CorporateAction) IsBasic() bool {
	return utils.Contains([]string{constants.Grouping, constants.Unfolding, constants.Update}, h.Description)
}

func (h CorporateAction) IsCashProceeds() bool {
	return utils.Contains([]string{constants.InterestOnEquity, constants.Dividend, constants.Income}, h.Description)
}

func (h CorporateAction) IsBonusProceeds() bool {
	return utils.Contains([]string{constants.Bonus}, h.Description)
}

func (CorporateAction) TableName() string {
	return "financial.corporate_actions"
}

type CorporateActionList struct {
	CorporateActionList []CorporateAction
}

type CorporateAction struct {
	Symbol           string    `gorm:"column:ticker"`
	Description      string    `gorm:"column:description"`
	Value            float64   `gorm:"column:value"`
	PaymentDate      time.Time `gorm:"column:payment_date"`
	ComDate          time.Time `gorm:"column:com_date"`
	TargetTicker     string    `gorm:"column:target_ticker"`
	CalculatedFactor float64   `gorm:"column:calculated_factor"`
	InitialDate      time.Time `gorm:"column:initial_date"`
}

func (t CorporateActionList) ToDomain() *[]domain.CorporateActions {
	var corporateActionsMap = map[string][]CorporateAction{}
	for _, value := range t.CorporateActionList {
		corporateActionsMap[value.Symbol] = append(corporateActionsMap[value.Symbol], value)
	}

	corporateActionDetail := []domain.CorporateActionDetail{}
	corporateActions := []domain.CorporateActions{}
	for k, v := range corporateActionsMap {
		for _, item := range v {
			corporateActionDetail = append(corporateActionDetail,
				domain.CorporateActionDetail{}.Create(item.Symbol, item.Description, item.Value, item.PaymentDate, item.ComDate, item.TargetTicker, item.CalculatedFactor, item.InitialDate))
		}
		corporateActions = append(corporateActions, domain.CorporateActions{}.Create(k, corporateActionDetail))
	}
	return &corporateActions
}
