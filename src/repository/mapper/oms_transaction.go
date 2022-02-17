package mapper

import "time"

func (OMSTransaction) TableName() string {
	return "wallet.oms_transactions"
}

type OMSTransaction struct {
	//Dados de transações
	ID        int       `gorm:"column:id"`
	Symbol    string    `gorm:"column:symbol"`
	Quantity  int       `gorm:"column:quantity"`
	Price     float64   `gorm:"column:price"`
	TradeDate time.Time `gorm:"column:trade_date"`

	//Dados calculados com base nos eventos corporativos
	PostEventQuantity int     `gorm:"column:post_event_quantity"`
	PostEventPrice    float64 `gorm:"column:post_event_price"`

	// Dados do evento corporativo que estão na Financial
	PostEventSymbol string    `gorm:"column:post_event_symbol"`
	EventFactor     float64   `gorm:"column:event_factor"`
	EventDate       time.Time `gorm:"column:event_date"`
	EventName       string    `gorm:"column:event_name"`
}
