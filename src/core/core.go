package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/constants"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/cei"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

var wg sync.WaitGroup

func Run() {
	start := time.Now()
	doApplyEvents(constants.AllCustomers)
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

var CorporateActionsAsc map[string][]mapper.CorporateAction
var CorporateActionsDesc map[string][]mapper.CorporateAction

var OMSCustomers []mapper.Customer
var ManualCustomers []mapper.Customer
var CEICustomers []mapper.Customer

var OMSSymbols []mapper.Symbol
var ManualSymbols []mapper.Symbol
var CEISymbols []mapper.Symbol

func doApplyEvents(customerCode string) {
	CorporateActionsAsc = repository.GetCorporateActions("asc")
	CorporateActionsDesc = repository.GetCorporateActions("desc")

	if customerCode == constants.AllCustomers {
		OMSCustomers = repository.GetOMSCustomers()
		ManualCustomers = repository.GetManualCustomers()
		CEICustomers = repository.GetCEICustomers()
	} else {
		Customer := mapper.Customer{}
		Customer.CustomerCode = customerCode
		OMSCustomers = append(OMSCustomers, Customer)
		ManualCustomers = append(ManualCustomers, Customer)
		CEICustomers = append(CEICustomers, Customer)
	}

	OMSSymbols = repository.GetSymbols(OMSCustomers)
	CEISymbols = repository.GetSymbols(CEICustomers)

	wg.Add(5)
	go doBasicOMSEvents()
	go doBasicManualEvents()
	go doBasicCEIEvents()
	go doProceedsOMSEvents()
	go doProceedsCEIEvents()
	wg.Wait()
}

func doBasicOMSEvents() {
	defer wg.Done()
	oms.BasicOMSEvents(OMSCustomers, CorporateActionsDesc)
}

func doBasicManualEvents() {
	defer wg.Done()
	manual.BasicManualEvents(ManualCustomers, CorporateActionsDesc)
}

func doBasicCEIEvents() {
	defer wg.Done()
	cei.BasicCEIEvents(CEICustomers, CorporateActionsDesc)
}

func doProceedsOMSEvents() {
	defer wg.Done()
	oms.ProceedsOMSEvents(CorporateActionsAsc, OMSCustomers, OMSSymbols)
}

func doProceedsCEIEvents() {
	defer wg.Done()
	cei.ProceedsCEIEvents(CorporateActionsAsc, CEICustomers, CEISymbols)
}
