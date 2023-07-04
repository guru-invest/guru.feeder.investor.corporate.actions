package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
	"github.com/sirupsen/logrus"
)

func init() {
	options.OPTIONS.Load()
}

func main() {
	//time_zone, _ := time.LoadLocation("America/Sao_Paulo")
	fmt.Println("Fluxo Iniciado")
	//core.ApplyEvents("fzVzgo8b")
	core.ApplyEvents(constants.AllCustomers)
	//core.ApplyEventsAfterInvestorSync("fzVzgo8b")
	defer func() {
		if r := recover(); r != nil {
			logrus.WithFields(logrus.Fields{
				"Service": "guru.feeder.investor.corporate.actions",
				"Caller":  "main",
				"Error":   r,
			}).Error("Erro inexperado no fluxo - [Panic]")

		}
	}()
	//core.Run()
	fmt.Println("Fluxo Finalizado")

	listener := make(chan os.Signal, 1)
	signal.Notify(listener, os.Interrupt, syscall.SIGTERM)
	<-listener
}
