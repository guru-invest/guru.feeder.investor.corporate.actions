package core

import (
	"fmt"
	"log"
	"time"

	"github.com/guru-invest/guru.corporate.actions/src/core/events"
	"github.com/guru-invest/guru.corporate.actions/src/repository"
	"github.com/guru-invest/guru.corporate.actions/src/repository/mapper"
)

func Run() {
	start := time.Now()
	doBasicEvents()
	elapsed := time.Since(start)
	fmt.Printf("Processs took %s\n", elapsed)
}

func doBasicEvents() {
	db := repository.SymbolRepository{}
	symbols, err := db.GetSymbols()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(symbols)

	events.Basic(mapper.OMSTransaction{})
}
