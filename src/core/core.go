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
	events.Basic(mapper.OMSTransaction{})
}
