package events

import (
	"encoding/json"
	"fmt"

	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func ApplyCorporateAction(OMSTransaction mapper.OMSTransaction) mapper.OMSTransaction {

	// TODO - Remover prints
	x, _ := json.Marshal(OMSTransaction)
	fmt.Println("Como era:")
	fmt.Println(string(x))

	// Quando for um evento de Atualização, o Fator deve ser 1, pois a Quantidade e o Preço não podem ser alterados.
	if OMSTransaction.EventName == singleton.New().Update {
		OMSTransaction.EventFactor = 1
	}

	OMSTransaction.PostEventQuantity = int(float64(OMSTransaction.Quantity) / OMSTransaction.EventFactor)
	OMSTransaction.PostEventPrice = utils.Truncate(OMSTransaction.Price*OMSTransaction.EventFactor, 2)

	// TODO - Remover prints
	y, _ := json.Marshal(OMSTransaction)
	fmt.Println("Como ficou:")
	fmt.Println(string(y))

	return OMSTransaction
}
