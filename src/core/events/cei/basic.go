package cei

import (
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyBasicCorporateAction(CEITransaction mapper.CEITransaction, corporate_action mapper.CorporateAction) mapper.CEITransaction {

	targetTicket := corporate_action.TargetTicker
	if corporate_action.TargetTicker == "" {
		targetTicket = corporate_action.Symbol
	}

	CEITransaction.EventFactor = corporate_action.CalculatedFactor
	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if corporate_action.Description == constants.Update {
		CEITransaction.EventFactor = 1
	}
	CEITransaction.EventName = corporate_action.Description
	CEITransaction.PostEventSymbol = targetTicket
	CEITransaction.EventDate = corporate_action.ComDate
	CEITransaction.PostEventQuantity = float64(CEITransaction.Quantity) / CEITransaction.EventFactor
	CEITransaction.PostEventPrice = utils.Truncate(CEITransaction.Amount/CEITransaction.PostEventQuantity, 2)
	CEITransaction.UpdatedAt = time.Now()

	// Processo cumulativo
	// CEITransaction.Quantity = utils.Truncate(CEITransaction.PostEventQuantity, 0)
	// CEITransaction.Price = CEITransaction.PostEventPrice

	return CEITransaction
}
