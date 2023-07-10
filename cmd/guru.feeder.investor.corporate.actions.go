package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/constants"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func init() {
	options.OPTIONS.Load()
}

func main() {
	br, _ := time.LoadLocation("America/Sao_Paulo")
	c := cron.New(cron.WithLocation(br))

	logrus.WithFields(logrus.Fields{
		"Service": "guru.feeder.investor.corporate.actions",
		"Caller":  "main",
	}).Info("Service Started")

	c.AddFunc("15 4 * * *", func() {
		start := time.Now()
		logrus.WithFields(logrus.Fields{
			"Service": "guru.feeder.investor.corporate.actions",
			"Caller":  "main",
		}).Info("Fluxo Iniciado")

		//core.ApplyEvents("fzVzgo8b")
		core.ApplyEvents(constants.AllCustomers)
		//core.ApplyEventsAfterInvestorSync("fzVzgo8b")

		logrus.WithFields(logrus.Fields{
			"Service": "guru.feeder.investor.corporate.actions",
			"Caller":  "main",
			"Elapsed": time.Since(start),
		}).Info("Fluxo Finalizado")
	})

	go c.Start()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}
