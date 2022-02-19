package core

import (
	"fmt"
	"log"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
)

func Run() {
	start := time.Now()
	doWork()
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)

	// Esse sleep fica aqui para segura os ultimos updates acumulados
	time.Sleep(time.Minute)
}

func doWork() {

	Symbols := repository.GetSymbols()
	totalOfSymbols := len(Symbols)
	currentSymbol := 0

	for _, value := range Symbols {

		log.Printf("%d de %d Symbols foram analisados\n", currentSymbol, totalOfSymbols)

		finished := make(chan bool)

		go doBasicOMSEvents(value.Name, finished)
		go doBasicManualEvents(value.Name, finished)

		<-finished
		currentSymbol += 1

	}

}

func doBasicOMSEvents(symbol string, finished chan bool) {
	CorporateActions := repository.GetCorporateActions(symbol)
	for _, value2 := range CorporateActions {
		symbol := symbol

		OMSTransaction := repository.GetOMSTransaction(symbol)

		for index, value3 := range OMSTransaction {

			if value3.TradeDate.After(value2.ComDate) {
				value3.EventName = "PADRAO"
				value3.PostEventSymbol = value3.Symbol
				value3.EventFactor = 1
				value3.EventDate, _ = time.Parse("2006-01-02", "2001-01-01")
				OMSTransaction[index] = oms.ApplyCorporateAction(value3)

			} else {

				value3.EventName = value2.Description
				value3.PostEventSymbol = value2.TargetTicker
				value3.EventFactor = value2.CalculatedFactor
				value3.EventDate = value2.ComDate
			}

			OMSTransaction[index] = oms.ApplyCorporateAction(value3)

		}

		go repository.UpdateOMSTransaction(OMSTransaction)

	}
	finished <- true
}

func doBasicManualEvents(symbol string, finished chan bool) {
	CorporateActions := repository.GetCorporateActions(symbol)
	for _, value2 := range CorporateActions {
		symbol := symbol

		ManualTransaction := repository.GetManualTransaction(symbol)

		for index, value3 := range ManualTransaction {

			if value3.TradeDate.After(value2.ComDate) {
				value3.EventName = "PADRAO"
				value3.PostEventSymbol = value3.Symbol
				value3.EventFactor = 1
				value3.EventDate, _ = time.Parse("2006-01-02", "2001-01-01")
				ManualTransaction[index] = manual.ApplyCorporateAction(value3)

			} else {

				value3.EventName = value2.Description
				value3.PostEventSymbol = value2.TargetTicker
				value3.EventFactor = value2.CalculatedFactor
				value3.EventDate = value2.ComDate
			}

			ManualTransaction[index] = manual.ApplyCorporateAction(value3)

		}

		go repository.UpdateManualTransaction(ManualTransaction)

	}
	finished <- true
}
