package oms

import (
	"crypto/sha1"
	"fmt"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

func ApplyProceedsCorporateAction(customer, symbol string, transactions map[string][]mapper.OMSTransaction, corporateActions map[string][]mapper.CorporateAction) []mapper.OMSProceeds {
	var result = []mapper.OMSProceeds{}

	for _, corporate_action := range corporateActions[symbol] {
		transaction_by_broker := map[float64]mapper.OMSProceeds{}

		for _, transaction := range transactions[customer] {
			if _, ok := transaction_by_broker[transaction.BrokerID]; !ok {
				transaction_by_broker[transaction.BrokerID] = mapper.OMSProceeds{}
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

			if broker != constants.Ideal {
				continue
			}

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

func applyCashProceeds(customer, symbol string, transaction_by_broker map[float64]mapper.OMSProceeds, transaction mapper.OMSTransaction, corporate_action mapper.CorporateAction) mapper.OMSProceeds {

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
	return mapper.OMSProceeds{}

}

func applyBonusProceeds(customer, symbol string, transaction_by_broker map[float64]mapper.OMSProceeds, transaction mapper.OMSTransaction, corporate_action mapper.CorporateAction) mapper.OMSProceeds {

	if transaction.BrokerID != constants.Ideal {
		return mapper.OMSProceeds{}
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

	return mapper.OMSProceeds{}

}
