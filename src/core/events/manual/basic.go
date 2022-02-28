package manual

import (
	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyBasicCorporateAction(manualTransaction mapper.ManualTransaction, corporate_action mapper.CorporateAction) mapper.ManualTransaction {

	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	manualTransaction.EventFactor = corporate_action.CalculatedFactor
	if corporate_action.Description == constants.Update {
		manualTransaction.EventFactor = 1
	}
	manualTransaction.EventName = corporate_action.Description
	manualTransaction.PostEventSymbol = corporate_action.TargetTicker
	manualTransaction.EventDate = corporate_action.ComDate
	manualTransaction.PostEventQuantity = float64(manualTransaction.Quantity) / manualTransaction.EventFactor
	manualTransaction.PostEventPrice = utils.Truncate(manualTransaction.Price*manualTransaction.EventFactor, 2)

	return manualTransaction
}
