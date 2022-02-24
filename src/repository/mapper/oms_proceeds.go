package mapper

import "time"

func (OMSProceeds) TableName() string {
	return "wallet.oms_proventos"
}

type OMSProceeds struct {
	ID           string    `gorm:"column:id"`
	CustomerCode string    `gorm:"column:customer_code"`
	BrokerID     float64   `gorm:"column:broker_id"`
	Symbol       string    `gorm:"column:symbol"`
	Quantity     float64   `gorm:"column:quantity"`
	Value        float64   `gorm:"column:value"`
	Amount       float64   `gorm:"column:amount"`
	Date         time.Time `gorm:"column:date"`
	Event        string    `gorm:"column:event"`
}
