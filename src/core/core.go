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
var Customers []mapper.Customer
var Symbols []mapper.Symbol

func doApplyEvents(customerCode string) {
	CorporateActionsAsc = repository.GetCorporateActions("asc")
	CorporateActionsDesc = repository.GetCorporateActions("desc")

	if customerCode == constants.AllCustomers {
		Customers = repository.GetCustomers()
	} else {
		Customer := mapper.Customer{}
		Customer.CustomerCode = customerCode
		Customers = append(Customers, Customer)
	}

	Symbols = repository.GetSymbols(Customers)

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
	oms.BasicOMSEvents(Customers, CorporateActionsDesc)
}

func doBasicManualEvents() {
	defer wg.Done()
	manual.BasicManualEvents(Customers, CorporateActionsDesc)
}

func doBasicCEIEvents() {
	defer wg.Done()
	cei.BasicCEIEvents(Customers, CorporateActionsDesc)
}

func doProceedsOMSEvents() {
	defer wg.Done()
	oms.ProceedsOMSEvents(CorporateActionsAsc, Customers, Symbols)
}

func doProceedsCEIEvents() {
	defer wg.Done()
	cei.ProceedsCEIEvents(CorporateActionsAsc, Customers, Symbols)
}
