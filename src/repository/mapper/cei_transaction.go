package mapper

import "time"

func (CEITransaction) TableName() string {
	return "wallet.investor_transactions"
}

type CEITransaction struct {
	// Dados de transações
	ID           int       `gorm:"column:id"`
	CustomerCode string    `gorm:"column:customer_code"`
	Symbol       string    `gorm:"column:symbol"`
	BrokerID     float64   `gorm:"column:broker_id"`
	Quantity     float64   `gorm:"column:quantity"`
	Price        float64   `gorm:"column:price"`
	TradeDate    time.Time `gorm:"column:trade_date"`
	Amount       float64   `gorm:"column:amount"`
	Side         int       `gorm:"column:side"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`

	// Dados calculados com base nos eventos corporativos
	PostEventQuantity float64 `gorm:"column:post_event_quantity"`
	PostEventPrice    float64 `gorm:"column:post_event_price"`

	// Dados do evento corporativo que estão na Financial
	PostEventSymbol string    `gorm:"column:post_event_symbol"`
	EventFactor     float64   `gorm:"column:event_factor"`
	EventDate       time.Time `gorm:"column:event_date"`
	EventName       string    `gorm:"column:event_name"`
}
