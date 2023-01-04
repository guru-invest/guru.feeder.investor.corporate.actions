package mapper

import (
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/utils"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/shopspring/decimal"
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

type CorporateAction struct {
	Symbol           string          `gorm:"column:ticker"`
	Description      string          `gorm:"column:description"`
	Value            decimal.Decimal `gorm:"column:value"`
	PaymentDate      time.Time       `gorm:"column:payment_date"`
	ComDate          time.Time       `gorm:"column:com_date"`
	TargetTicker     string          `gorm:"column:target_ticker"`
	CalculatedFactor decimal.Decimal `gorm:"column:calculated_factor"`
	InitialDate      time.Time       `gorm:"column:initial_date"`
}

func (t CorporateAction) ToDomain(mapper []CorporateAction) *domain.CorporateActionList {
	list := domain.CorporateActionList{}

	for _, item := range mapper {
		list.Map[item.Symbol] = append(list.Map[item.Symbol],
			domain.CorporateAction{
				Symbol:           item.Symbol,
				Description:      item.Description,
				Value:            item.Value,
				PaymentDate:      item.PaymentDate,
				ComDate:          item.ComDate,
				TargetTicker:     item.TargetTicker,
				CalculatedFactor: item.CalculatedFactor,
				InitialDate:      item.InitialDate,
			})
	}

	return &list
}
