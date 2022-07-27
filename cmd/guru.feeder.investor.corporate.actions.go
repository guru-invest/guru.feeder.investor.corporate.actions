package main

import (
	"fmt"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
)

func init() {
	options.OPTIONS.Load()
}

func main() {
	//time_zone, _ := time.LoadLocation("America/Sao_Paulo")
	fmt.Println("Fluxo Iniciado")
	//core.ApplyEvents("fzVzgo8b")
	//core.ApplyEventsAfterInvestorSync("fzVzgo8b")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic: %v \n", r)

		}
	}()
	core.Run()
	fmt.Println("Fluxo Finalizado")
	// c := cron.New(cron.WithLocation(time_zone))
	// c.AddFunc("30 2 * * *", func() { core.Run() })
	// go c.Start()
	// sig := make(chan os.Signal, 1)
	// signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	// <-sig
}
