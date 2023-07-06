package mapper

import "time"

func (ManualTransaction) TableName() string {
	return "wallet.manual_transactions"
}

type ManualTransaction struct {
	// Dados de transações
	ID             int       `gorm:"column:id"`
	CustomerCode   string    `gorm:"column:customer_code"`
	BrokerID       float64   `gorm:"column:broker_id"`
	InvestmentType int       `gorm:"column:investment_type"`
	Symbol         string    `gorm:"column:symbol"`
	Quantity       float64   `gorm:"column:quantity"`
	Price          float64   `gorm:"column:price"`
	Amount         float64   `gorm:"column:amount"`
	Side           int       `gorm:"column:side"`
	TradeDate      time.Time `gorm:"column:trade_date"`
	SourceType     string    `gorm:"column:source_type"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`

	// Dados calculados com base nos eventos corporativos
	PostEventQuantity float64 `gorm:"column:post_event_quantity"`
	PostEventPrice    float64 `gorm:"column:post_event_price"`

	// Dados do evento corporativo que estão na Financial
	PostEventSymbol string    `gorm:"column:post_event_symbol"`
	EventFactor     float64   `gorm:"column:event_factor"`
	EventDate       time.Time `gorm:"column:event_date"`
	EventName       string    `gorm:"column:event_name"`

	Hash_ID string `gorm:"column:hash_id"`
}
