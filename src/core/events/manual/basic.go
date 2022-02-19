package manual

import (
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyCorporateAction(ManualTransaction mapper.ManualTransaction) mapper.ManualTransaction {

	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if ManualTransaction.EventName == singleton.New().Update {
		ManualTransaction.EventFactor = 1
	}

	ManualTransaction.PostEventQuantity = float64(ManualTransaction.Quantity) / ManualTransaction.EventFactor
	ManualTransaction.PostEventPrice = utils.Truncate(ManualTransaction.Price*ManualTransaction.EventFactor, 2)

	return ManualTransaction
}
