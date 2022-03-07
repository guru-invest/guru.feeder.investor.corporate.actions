package cei

import (
	"crypto/sha1"
	"fmt"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

func ApplyProceedsCorporateAction(customer, symbol string, transactions map[string][]mapper.CEITransaction, corporateActions map[string][]mapper.CorporateAction) []mapper.CEIProceeds {
	var result = []mapper.CEIProceeds{}

	for _, corporate_action := range corporateActions[symbol] {
		transaction_by_broker := map[float64]mapper.CEIProceeds{}

		for _, transaction := range transactions[customer] {
			if transaction.BrokerID == constants.Ideal {
				continue
			}

			if _, ok := transaction_by_broker[transaction.BrokerID]; !ok {
				transaction_by_broker[transaction.BrokerID] = mapper.CEIProceeds{}
			}

			if transaction.Symbol != symbol {
				continue
			}

			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			if corporate_action.IsCashProceeds() {
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

					StringID := fmt.Sprintf("%s %f %s %f %f %f %s %s %s",
						entry.CustomerCode,
						entry.BrokerID,
						entry.Symbol,
						entry.Quantity,
						entry.Value,
						entry.Amount,
						entry.Date.String(),
						entry.Event,
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
		entry.Quantity += float64(transaction.Quantity)
		entry.Value = corporate_action.Value
		entry.Amount = entry.Quantity * entry.Value
		entry.Date = corporate_action.ComDate
		entry.Event = corporate_action.Description
		return entry
	}
	return mapper.CEIProceeds{}

}

func applyBonusProceeds(customer, symbol string, transaction_by_broker map[float64]mapper.CEIProceeds, transaction mapper.CEITransaction, corporate_action mapper.CorporateAction) mapper.CEIProceeds {

	if transaction.BrokerID == constants.Ideal {
		return mapper.CEIProceeds{}
	}

	if entry, ok := transaction_by_broker[transaction.BrokerID]; ok {
		entry.CustomerCode = customer
		entry.Symbol = symbol
		entry.Quantity += float64(transaction.Quantity)
		entry.Value = corporate_action.Value
		entry.Amount = utils.Truncate((entry.Quantity / entry.Value), 0)
		entry.Date = corporate_action.ComDate
		entry.Event = corporate_action.Description
		return entry
	}

	return mapper.CEIProceeds{}

}
