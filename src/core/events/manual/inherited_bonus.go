package manual

import (
	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

func ApplyInheritedBonusAction(manualTransaction mapper.ManualTransaction, proceed mapper.OMSProceeds) mapper.ManualTransaction {
	if proceed.Event == constants.Bonus && proceed.Quantity > 0 {
		manualTransaction.CustomerCode = proceed.CustomerCode
		manualTransaction.BrokerID = proceed.BrokerID
		manualTransaction.InvestmentType = constants.BonusInvestmentType
		manualTransaction.Symbol = proceed.Symbol
		manualTransaction.Quantity = proceed.Amount
		manualTransaction.Price = constants.MinimalPrice
		manualTransaction.Amount = manualTransaction.Quantity * manualTransaction.Price
		manualTransaction.Side = constants.Purchase
		manualTransaction.TradeDate = proceed.Date // TODO com_date ou initial_date do evento ?
		manualTransaction.SourceType = constants.ManualSourceType
		manualTransaction.EventDate = proceed.Date
		manualTransaction.EventName = proceed.Event
	}
	return manualTransaction
}
