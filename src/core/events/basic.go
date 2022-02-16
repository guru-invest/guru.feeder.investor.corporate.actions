package events

import (
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func Basic(OMSTransaction mapper.OMSTransaction) mapper.OMSTransaction {

	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if OMSTransaction.Event == singleton.New().Update {
		OMSTransaction.Factor = 1
	}

	OMSTransaction.PostEventQuantity = int(float64(OMSTransaction.Quantity) / OMSTransaction.Factor)
	OMSTransaction.PostEventPrice = utils.Truncate(OMSTransaction.Price*OMSTransaction.Factor, 2)

	return OMSTransaction
}
