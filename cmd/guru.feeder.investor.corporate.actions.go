package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/core"
	"github.com/guru-invest/guru.feeder.investor.corporate.actions/src/crossCutting/options"
	"github.com/robfig/cron/v3"
)

func init() {
	options.OPTIONS.Load()
}

func main() {
	time_zone, _ := time.LoadLocation("America/Sao_Paulo")

	c := cron.New(cron.WithLocation(time_zone))
	c.AddFunc("30 2 * * *", func() { core.Run() })
	go c.Start()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig
}
