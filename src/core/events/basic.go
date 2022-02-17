package events

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/repository"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

// Eventos basicos contemplam Grupamento, Desdobramento e Atualização
func Basic(OMSTransaction mapper.OMSTransaction) mapper.OMSTransaction {

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

func GetSymbols() []mapper.Symbol {
	db := repository.SymbolRepository{}
	symbols, err := db.GetSymbols()
	if err != nil {
		log.Println(err)
		return []mapper.Symbol{}
	}

	return symbols
}

func GetCorporateActions(symbol string) []mapper.CorporateAction {
	db := repository.CorporateActionRepository{}
	corporate_actions, err := db.GetCorporateActions(symbol)
	if err != nil {
		log.Println(err)
		return []mapper.CorporateAction{}
	}

	return corporate_actions
}

func GetOMSTransaction(symbol, event string, begin_date, end_date time.Time) []mapper.OMSTransaction {
	db := repository.OMSTransactionRepository{}
	oms_transaction, err := db.GetOMSTransactions(symbol, event, begin_date, end_date)
	if err != nil {
		log.Println(err)
		return []mapper.OMSTransaction{}
	}

	return oms_transaction
}
