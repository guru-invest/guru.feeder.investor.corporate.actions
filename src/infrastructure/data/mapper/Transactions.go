package mapper

import (
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/domain"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	// Dados de transações
	ID           int             `gorm:"column:id"`
	CustomerCode string          `gorm:"column:customer_code"`
	BrokerID     int64           `gorm:"column:broker_id"`
	Symbol       string          `gorm:"column:symbol"`
	Quantity     decimal.Decimal `gorm:"column:quantity"`
	Price        decimal.Decimal `gorm:"column:price"`
	TradeDate    time.Time       `gorm:"column:trade_date"`
	Amount       decimal.Decimal `gorm:"column:amount"`
	Side         int             `gorm:"column:side"`

	// Dados calculados com base nos eventos corporativos
	PostEventQuantity decimal.Decimal `gorm:"column:post_event_quantity"`
	PostEventPrice    decimal.Decimal `gorm:"column:post_event_price"`

	// Dados do evento corporativo que estão na Financial
	PostEventSymbol string          `gorm:"column:post_event_symbol"`
	EventFactor     decimal.Decimal `gorm:"column:event_factor"`
	EventDate       time.Time       `gorm:"column:event_date"`
	EventName       string          `gorm:"column:event_name"`
}

func (t Transaction) ToDomain(mapper []Transaction) *domain.TransactionList {
	list := domain.TransactionList{}

	for _, item := range mapper {
		list.Map[item.CustomerCode] = append(list.Map[item.CustomerCode],
			domain.Transaction{
				ID:                item.ID,
				CustomerCode:      item.CustomerCode,
				BrokerID:          item.BrokerID,
				Symbol:            item.Symbol,
				Quantity:          item.Quantity,
				Price:             item.Price,
				TradeDate:         item.TradeDate,
				Amount:            item.Amount,
				Side:              item.Side,
				PostEventQuantity: item.PostEventQuantity,
				PostEventPrice:    item.PostEventPrice,
				PostEventSymbol:   item.PostEventSymbol,
				EventFactor:       item.EventFactor,
				EventDate:         item.EventDate,
				EventName:         item.EventName,
			})
	}

	return &list
}
