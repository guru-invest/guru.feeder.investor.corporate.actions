package core

import (
	"fmt"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/core/events"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

func Run() {
	start := time.Now()
	doBasicEvents()
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

func doBasicEvents() {

	Symbols := events.GetSymbols()
	fmt.Println(Symbols)

	CorporateActions := events.GetCorporateActions("BIDI11")
	fmt.Println(CorporateActions)

	events.Basic(mapper.OMSTransaction{})
}
