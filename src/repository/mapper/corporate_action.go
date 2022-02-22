package mapper

import (
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

func (h CorporateAction) IsBasic() bool {
	return utils.Contains([]string{constants.Grouping, constants.Unfolding, constants.Update}, h.Description)
}

func (CorporateAction) TableName() string {
	return "financial.corporate_actions"
}

type CorporateAction struct {
	Symbol           string    `gorm:"column:ticker"`
	Description      string    `gorm:"column:description"`
	ComDate          time.Time `gorm:"column:com_date"`
	TargetTicker     string    `gorm:"column:target_ticker"`
	CalculatedFactor float64   `gorm:"column:calculated_factor"`
}
