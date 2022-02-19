package cei

import (
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyCorporateAction(CEITransaction mapper.CEITransaction) mapper.CEITransaction {

	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if CEITransaction.EventName == singleton.New().Update {
		CEITransaction.EventFactor = 1
	}

	CEITransaction.PostEventQuantity = float64(CEITransaction.Quantity) / CEITransaction.EventFactor
	CEITransaction.PostEventPrice = utils.Truncate(CEITransaction.Price*CEITransaction.EventFactor, 2)

	return CEITransaction
}
