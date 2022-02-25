package manual

import (
	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyBasicCorporateAction(manualTransaction mapper.ManualTransaction) mapper.ManualTransaction {

	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if manualTransaction.EventName == constants.Update {
		manualTransaction.EventFactor = 1
	}

	manualTransaction.PostEventQuantity = float64(manualTransaction.Quantity) / manualTransaction.EventFactor
	manualTransaction.PostEventPrice = utils.Truncate(manualTransaction.Price*manualTransaction.EventFactor, 2)

	return manualTransaction
}
