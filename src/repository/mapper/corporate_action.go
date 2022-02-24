package mapper

import (
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

func (h CorporateAction) IsBasic() bool {
	return utils.Contains([]string{constants.Grouping, constants.Unfolding, constants.Update}, h.Description)
}

func (h CorporateAction) IsCashProceeds() bool {
	return utils.Contains([]string{constants.InterestOnEquity, constants.Dividend, constants.Income}, h.Description)
}

func (CorporateAction) TableName() string {
	return "financial.corporate_actions"
}

type CorporateAction struct {
	Symbol           string    `gorm:"column:ticker"`
	Description      string    `gorm:"column:description"`
	Value            float64   `gorm:"column:value"`
	PaymentDate      time.Time `gorm:"column:payment_date"`
	ComDate          time.Time `gorm:"column:com_date"`
	TargetTicker     string    `gorm:"column:target_ticker"`
	CalculatedFactor float64   `gorm:"column:calculated_factor"`
}
