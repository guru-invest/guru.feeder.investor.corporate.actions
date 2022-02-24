package cei

import (
	"crypto/sha1"
	"fmt"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

func ApplyCashProceedsCorporateAction(Customer, Symbol string, Transactions map[string][]mapper.CEITransaction, CorporateActions map[string][]mapper.CorporateAction) []mapper.CEIProceeds {

	var result = []mapper.CEIProceeds{}
	for _, corporate_action := range CorporateActions[Symbol] {
		var partial_result = mapper.CEIProceeds{}
		for _, transaction := range Transactions[Customer] {

			if transaction.Symbol != Symbol {
				continue
			}

			if transaction.TradeDate.After(corporate_action.ComDate) {
				continue
			}

			if corporate_action.IsCashProceeds() {
				partial_result.CustomerCode = Customer
				partial_result.BrokerID = transaction.BrokerID
				partial_result.Symbol = Symbol
				partial_result.Quantity += float64(transaction.Quantity)
				partial_result.Value = corporate_action.Value
				partial_result.Amount = partial_result.Quantity * partial_result.Value
				partial_result.Date = corporate_action.ComDate
				partial_result.Event = corporate_action.Description
			}

		}
		if partial_result.Quantity > 0 {
			StringID := fmt.Sprintf("%s %f %s %f %f %f %s %s %s",
				partial_result.CustomerCode,
				partial_result.BrokerID,
				partial_result.Symbol,
				partial_result.Quantity,
				partial_result.Value,
				partial_result.Amount,
				partial_result.Date.String(),
				partial_result.Event,
				corporate_action.PaymentDate.String())

			HashID := sha1.New()
			HashID.Write([]byte(StringID))

			partial_result.ID = fmt.Sprintf("%x", HashID.Sum(nil))
			result = append(result, partial_result)
		}
	}

	return result
}
