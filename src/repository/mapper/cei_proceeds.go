package mapper

import "time"

func (CEIProceeds) TableName() string {
	return "wallet.cei_proventos"
}

type CEIProceeds struct {
	ID           string    `gorm:"column:id"`
	CustomerCode string    `gorm:"column:customer_code"`
	BrokerID     float64   `gorm:"column:broker_id"`
	Symbol       string    `gorm:"column:symbol"`
	Quantity     float64   `gorm:"column:quantity"`
	Value        float64   `gorm:"column:value"`
	Amount       float64   `gorm:"column:amount"`
	Event        string    `gorm:"column:event"`
	InitialDate  time.Time `gorm:"column:initial_date"`
	ComDate      time.Time `gorm:"column:com_date"`
	PaymentDate  time.Time `gorm:"column:payment_date"`
}
