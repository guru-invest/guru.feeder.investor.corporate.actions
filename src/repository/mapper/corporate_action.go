package mapper

import "time"

func (CorporateAction) TableName() string {
	return "financial.corporate_actions"
}

type CorporateAction struct {
	Description      string    `gorm:"column:description"`
	ComDate          time.Time `gorm:"column:com_date"`
	TargetTicker     string    `gorm:"column:target_ticker"`
	CalculatedFactor float64   `gorm:"column:calculated_factor"`
}
