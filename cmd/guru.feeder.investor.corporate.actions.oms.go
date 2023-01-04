package main

import (
	"fmt"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/app/handlers"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions.oms/src/crossCutting/options"
)

func init() {
	options.OPTIONS.Load()
}

func main() {
	fmt.Println("Fluxo Iniciado")
	handlers.ApplyEventsHandler{}.NewApplyEventsHandler("fzVzgo8b")
	//core.ApplyEventsAfterInvestorSync("fzVzgo8b")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %v \n", r)

		}
	}()
	fmt.Println("Fluxo Finalizado")
}
