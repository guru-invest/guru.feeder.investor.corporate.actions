package cei

import (
	"crypto/sha1"
	"fmt"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/utils"
)

func ApplyProceedsCorporateAction(customer, symbol string, transactions map[string][]mapper.CEITransaction, corporateActions map[string][]mapper.CorporateAction) []mapper.CEIProceeds {
	var result = []mapper.CEIProceeds{}

	for _, corporate_action := range corporateActions[symbol] {
		transaction_by_broker := map[float64]mapper.CEIProceeds{}

		for _, transaction := range transactions[customer] {

			if transaction.BrokerID == constants.Genial {
				continue
			}

			if _, ok := transaction_by_broker[transaction.BrokerID]; !ok {
				transaction_by_broker[transaction.BrokerID] = mapper.CEIProceeds{}
			}

			if transaction.Symbol != symbol {
				continue
			}

			if transaction.TradeDate.After(corporate_action.InitialDate) && !transaction.TradeDate.Equal(corporate_action.InitialDate) {
				continue
			}

			if corporate_action.IsCashProceeds() {

				if corporate_action.PaymentDate.Year() == 0 {
					continue
				}

				transaction_by_broker[transaction.BrokerID] =
					applyCashProceeds(
						customer,
						symbol,
						transaction_by_broker,
						transaction,
						corporate_action,
					)
				continue
			}

			if corporate_action.IsBonusProceeds() {
				transaction_by_broker[transaction.BrokerID] =
					applyBonusProceeds(
						customer,
						symbol,
						transaction_by_broker,
						transaction,
						corporate_action,
					)
				continue
			}

		}

		for broker := range transaction_by_broker {

			if entry, ok := transaction_by_broker[broker]; ok {
				if entry.Quantity > 0 {

					StringID := fmt.Sprintf("%s %f %s %f %f %f %s %s %s %s %s",
						entry.CustomerCode,
						entry.BrokerID,
						entry.Symbol,
						entry.Quantity,
						entry.Value,
						entry.Amount,
						entry.Event,
						entry.InitialDate.String(),
						entry.ComDate.String(),
						entry.PaymentDate.String(),
						corporate_action.PaymentDate.String())

					HashID := sha1.New()
					HashID.Write([]byte(StringID))

					entry.ID = fmt.Sprintf("%x", HashID.Sum(nil))
					entry.BrokerID = broker
					result = append(result, entry)
				}
			}
		}
	}

	return result
}

func applyCashProceeds(customer, symbol string, transaction_by_broker map[float64]mapper.CEIProceeds, transaction mapper.CEITransaction, corporate_action mapper.CorporateAction) mapper.CEIProceeds {

	if entry, ok := transaction_by_broker[transaction.BrokerID]; ok {
		entry.CustomerCode = customer
		entry.Symbol = symbol

		if transaction.Side == constants.Sale && transaction.Quantity > 0 {
			transaction.Quantity = transaction.Quantity * -1
		}
		entry.Quantity += transaction.Quantity
		entry.Quantity = utils.Truncate(entry.Quantity, 0)

		entry.Value = corporate_action.Value

		entry.Amount = entry.Quantity * entry.Value
		if corporate_action.Description == constants.InterestOnEquity {
			entry.Amount = entry.Amount - (entry.Amount * constants.InterestOnEquityIRPercent)
		}

		entry.Event = corporate_action.Description
		entry.InitialDate = corporate_action.InitialDate
		entry.ComDate = corporate_action.ComDate
		entry.PaymentDate = corporate_action.PaymentDate
		return entry
	}
	return mapper.CEIProceeds{}

}

func applyBonusProceeds(customer, symbol string, transaction_by_broker map[float64]mapper.CEIProceeds, transaction mapper.CEITransaction, corporate_action mapper.CorporateAction) mapper.CEIProceeds {

	if transaction.BrokerID == constants.Genial {
		return mapper.CEIProceeds{}
	}

	if entry, ok := transaction_by_broker[transaction.BrokerID]; ok {
		entry.CustomerCode = customer
		entry.Symbol = symbol
		entry.Quantity += transaction.Quantity
		entry.Quantity = utils.Truncate(entry.Quantity, 0)
		entry.Value = corporate_action.Value
		entry.Amount = utils.Truncate((entry.Quantity / entry.Value), 0)
		entry.Event = corporate_action.Description
		entry.InitialDate = corporate_action.InitialDate
		entry.ComDate = corporate_action.ComDate
		entry.PaymentDate = corporate_action.PaymentDate
		return entry
	}

	return mapper.CEIProceeds{}

}
