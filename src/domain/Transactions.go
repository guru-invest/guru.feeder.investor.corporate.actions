package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type TransactionList struct {
	Map  map[string][]Transaction
	List []Transaction
}

type Transaction struct {
	ID                int
	CustomerCode      string
	BrokerID          int64
	Symbol            string
	Quantity          decimal.Decimal
	Price             decimal.Decimal
	TradeDate         time.Time
	Amount            decimal.Decimal
	Side              int
	PostEventQuantity decimal.Decimal
	PostEventPrice    decimal.Decimal
	PostEventSymbol   string
	EventFactor       decimal.Decimal
	EventDate         time.Time
	EventName         string
}
