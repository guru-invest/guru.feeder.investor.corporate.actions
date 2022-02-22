package core

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/cei"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
	"github.com/guru-invest/guru.corporate.actions/src/singleton"
	"github.com/guru-invest/guru.corporate.actions/src/utils"
)

func Run() {
	start := time.Now()
	doApplyBasicEvents()
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)

	// Esse sleep fica aqui para segura os ultimos updates acumulados
	time.Sleep(time.Minute)
}

func doApplyBasicEvents() {

	Symbols := repository.GetSymbols()
	totalOfSymbols := len(Symbols)
	currentSymbol := 0

	for _, value := range Symbols {

		log.Printf("%d de %d Symbols foram analisados\n", currentSymbol, totalOfSymbols)

		finished := make(chan bool)

		go doBasicOMSEvents(value.Name, finished)
		go doBasicManualEvents(value.Name, finished)
		go doBasicCEIEvents(value.Name, finished)

		<-finished
		currentSymbol += 1

	}

}

func doBasicOMSEvents(symbol string, finished chan bool) {
	CorporateActions := repository.GetCorporateActions(symbol)
	OMSTransactionBkp := []mapper.OMSTransaction{}

	for _, value2 := range CorporateActions {
		symbol := symbol

		OMSTransaction := repository.GetOMSTransaction(symbol)

		for index, value3 := range OMSTransaction {
			OMSTransactionBkp = append(OMSTransactionBkp, value3)

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if value3.TradeDate.After(value2.ComDate) {
				continue
			}

			// Se o Event name for de Atualização, Grupamento ou Desobramento, aplica eventos corporativos basicos
			if utils.Contains([]string{singleton.New().Grouping, singleton.New().Unfolding, singleton.New().Update}, value2.Description) {
				value3.EventName = value2.Description
				value3.PostEventSymbol = value2.TargetTicker
				value3.EventFactor = value2.CalculatedFactor
				value3.EventDate = value2.ComDate
				OMSTransaction[index] = oms.ApplyBasicCorporateAction(value3)
				continue
			}

		}

		if !cmp.Equal(OMSTransactionBkp, OMSTransaction) {
			go repository.UpdateOMSTransaction(OMSTransaction)
		}

	}
	finished <- true
}

func doBasicManualEvents(symbol string, finished chan bool) {
	CorporateActions := repository.GetCorporateActions(symbol)
	ManualTransactionBkp := []mapper.ManualTransaction{}

	for _, value2 := range CorporateActions {
		symbol := symbol

		ManualTransaction := repository.GetManualTransaction(symbol)

		for index, value3 := range ManualTransaction {
			ManualTransactionBkp = append(ManualTransactionBkp, value3)

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if value3.TradeDate.After(value2.ComDate) {
				continue
			}

			if utils.Contains([]string{singleton.New().Grouping, singleton.New().Unfolding, singleton.New().Update}, value2.Description) {
				value3.EventName = value2.Description
				value3.PostEventSymbol = value2.TargetTicker
				value3.EventFactor = value2.CalculatedFactor
				value3.EventDate = value2.ComDate
				ManualTransaction[index] = manual.ApplyBasicCorporateAction(value3)
				continue
			}

		}

		if !cmp.Equal(ManualTransactionBkp, ManualTransaction) {
			go repository.UpdateManualTransaction(ManualTransaction)
		}

	}
	finished <- true
}

func doBasicCEIEvents(symbol string, finished chan bool) {
	CorporateActions := repository.GetCorporateActions(symbol)
	CEITransactionBkp := []mapper.CEITransaction{}

	for _, value2 := range CorporateActions {
		symbol := symbol

		CEITransaction := repository.GetCEITransaction(symbol)

		for index, value3 := range CEITransaction {
			CEITransactionBkp = append(CEITransactionBkp, value3)

			// Se a data de com_date for maior, significa que eu não precios aplicar este evento nesta transação
			if value3.TradeDate.After(value2.ComDate) {
				continue

			}

			if utils.Contains([]string{singleton.New().Grouping, singleton.New().Unfolding, singleton.New().Update}, value2.Description) {
				value3.EventName = value2.Description
				value3.PostEventSymbol = value2.TargetTicker
				value3.EventFactor = value2.CalculatedFactor
				value3.EventDate = value2.ComDate
				CEITransaction[index] = cei.ApplyBasicCorporateAction(value3)
				continue
			}

		}

		if !cmp.Equal(CEITransactionBkp, CEITransaction) {
			go repository.UpdateCEITransaction(CEITransaction)
		}

	}
	finished <- true
}
