package manual

import (
	"crypto/sha1"
	"fmt"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyBasicCorporateAction(manualTransaction mapper.ManualTransaction, corporate_action mapper.CorporateAction) mapper.ManualTransaction {
	targetTicket := corporate_action.TargetTicker
	if corporate_action.TargetTicker == "" {
		targetTicket = corporate_action.Symbol
	}
	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	manualTransaction.EventFactor = corporate_action.CalculatedFactor
	if corporate_action.Description == constants.Update {
		manualTransaction.EventFactor = 1
	}
	manualTransaction.EventName = corporate_action.Description
	manualTransaction.PostEventSymbol = targetTicket
	manualTransaction.EventDate = corporate_action.ComDate
	manualTransaction.PostEventQuantity = float64(manualTransaction.Quantity) / manualTransaction.EventFactor
	manualTransaction.PostEventPrice = utils.Truncate(manualTransaction.Amount/manualTransaction.PostEventQuantity, 2)

	// Processo cumulativo
	// manualTransaction.Quantity = utils.Truncate(manualTransaction.PostEventQuantity, 0)
	// manualTransaction.Price = manualTransaction.PostEventPrice

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

	return manualTransaction
}
