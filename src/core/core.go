package core

import (
	"fmt"
	"log"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/core/events"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
)

func Run() {
	start := time.Now()
	doWork()
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

func doWork() {

	// Retorna todos os Symbols que não foram aplicados eventos.
	// Symbols := events.GetSymbols()
	// fmt.Println(Symbols)

	// Retorna todos os Eventos Corporativos de um determinado Symbol.
	// CorporateActions := events.GetCorporateActions("BIDI11")
	// fmt.Println(CorporateActions)

	// Retorna as transações que devem ser aplicados os eventos corporativos.
	// Aqui a data inicial é o com_date do meu proximo evento e a data final é o com_date do meu evento atual
	// Dessa forma eu consigo saber quais as transações realmente precisam sofrer a aplicação do evento
	// OMSTransaction := events.GetOMSTransaction("BIDI11", CorporateActions[0].Description, CorporateActions[1].ComDate, CorporateActions[0].ComDate)
	// fmt.Println(OMSTransaction)

	// events.Basic(OMSTransaction[0])

	Symbols := repository.GetSymbols()
	totalOfSymbols := len(Symbols)
	currentSymbol := 0

	for _, value := range Symbols {

		log.Printf("%d de %d Symbols foram analisados\n", currentSymbol, totalOfSymbols)
		doBasicEvents(value.Name)

		currentSymbol += 1

	}

}

func doBasicEvents(symbol string) {
	CorporateActions := repository.GetCorporateActions(symbol)
	for index2, value2 := range CorporateActions {

		var begin_date time.Time
		if index2 >= len(CorporateActions)-1 {
			begin_date = time.Now().AddDate(-5, 0, 0)
		} else {
			begin_date = CorporateActions[index2+1].ComDate
		}

		end_date := value2.ComDate
		event := value2.Description
		target_symbol := value2.TargetTicker
		symbol := symbol

		OMSTransaction := repository.GetOMSTransaction(symbol, target_symbol, event, begin_date, end_date)

		for _, value3 := range OMSTransaction {
			events.ApplyCorporateAction(value3)
			//TODO - persistir o new_oms_transaction
		}
	}

}
