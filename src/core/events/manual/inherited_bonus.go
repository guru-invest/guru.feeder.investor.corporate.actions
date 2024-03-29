package manual

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
)

func ApplyInheritedBonusActionCEI(manualTransaction mapper.ManualTransaction, proceed mapper.CEIProceeds) mapper.ManualTransaction {
	if proceed.Event == constants.Bonus && proceed.Quantity > 0 {
		manualTransaction.CustomerCode = proceed.CustomerCode
		manualTransaction.BrokerID = proceed.BrokerID
		manualTransaction.InvestmentType = constants.BonusInvestmentType
		manualTransaction.Symbol = proceed.Symbol
		manualTransaction.Quantity = proceed.Amount
		manualTransaction.Price = constants.MinimalValue
		manualTransaction.Amount = constants.MinimalValue
		manualTransaction.Side = constants.Purchase
		manualTransaction.TradeDate = proceed.InitialDate // TODO - Validar com TOM
		manualTransaction.SourceType = constants.BonusInvestorSourceType
		manualTransaction.EventDate = proceed.InitialDate
		manualTransaction.EventName = proceed.Event
		manualTransaction.UpdatedAt = time.Now()

		StringID := fmt.Sprintf("%s %f %d %s %f %f %f %d %s %s %s %s",
			manualTransaction.CustomerCode,
			manualTransaction.BrokerID,
			manualTransaction.InvestmentType,
			manualTransaction.Symbol,
			manualTransaction.Quantity,
			manualTransaction.Price,
			manualTransaction.Amount,
			manualTransaction.Side,
			manualTransaction.TradeDate.String(),
			manualTransaction.SourceType,
			manualTransaction.EventDate.String(),
			manualTransaction.EventName,
		)

		HashID := sha1.New()
		HashID.Write([]byte(StringID))

		manualTransaction.Hash_ID = fmt.Sprintf("%x", HashID.Sum(nil))
	}
	return manualTransaction
}
