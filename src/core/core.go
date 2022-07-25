package core

import (
	"fmt"
	"sync"
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/cei"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/manual"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core/events/oms"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/repository/mapper"
)

var wg sync.WaitGroup

func Run() {
	start := time.Now()
	ApplyEvents(constants.AllCustomers)
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

func ApplyEvents(customerCode string) {
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

	OMSSymbols = repository.GetOMSSymbols(OMSCustomers)
	CEISymbols = repository.GetCEISymbols(CEICustomers)

	wg.Add(5)
	go doBasicOMSEvents()
	go doBasicManualEvents()
	go doBasicCEIEvents(false)
	go doProceedsOMSEvents(false)
	go doProceedsCEIEvents(false)
	wg.Wait()

	go repository.NewWalletConnector().ResyncAveragePrice()

	fmt.Println("Finalizou")
}

func ApplyEventsAfterInvestorSync(customerCode string) error {
	options.OPTIONS.Load()
	CorporateActionsAsc = repository.GetCorporateActions("asc")
	CorporateActionsDesc = repository.GetCorporateActions("desc")

	Customer := mapper.Customer{
		CustomerCode: customerCode,
		CreatedAT:    time.Now().String(),
	}

	CEICustomers = []mapper.Customer{Customer}

	CEISymbols = repository.GetCEISymbols(CEICustomers)

	//wg.Add(2)
	doBasicCEIEventsServerless(true)
	doProceedsCEIEventsServerless(true)
	//wg.Wait()

	return nil
}

func doBasicOMSEvents() {
	defer wg.Done()
	oms.BasicOMSEvents(OMSCustomers, CorporateActionsDesc)
}

func doBasicManualEvents() {
	defer wg.Done()
	manual.BasicManualEvents(ManualCustomers, CorporateActionsDesc)
}

func doBasicCEIEvents(isStateLess bool) {
	defer wg.Done()
	cei.BasicCEIEvents(CEICustomers, CorporateActionsDesc, isStateLess)
}

func doProceedsOMSEvents(isStateLess bool) {
	defer wg.Done()
	oms.ProceedsOMSEvents(CorporateActionsAsc, OMSCustomers, OMSSymbols, isStateLess)
}

func doProceedsCEIEvents(isStateLess bool) {
	defer wg.Done()
	cei.ProceedsCEIEvents(CorporateActionsAsc, CEICustomers, CEISymbols, isStateLess)
}

func doBasicCEIEventsServerless(isStateLess bool) {
	cei.BasicCEIEvents(CEICustomers, CorporateActionsDesc, isStateLess)
}
func doProceedsCEIEventsServerless(isStateLess bool) {
	cei.ProceedsCEIEvents(CorporateActionsAsc, CEICustomers, CEISymbols, isStateLess)
}
