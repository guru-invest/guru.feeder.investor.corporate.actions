package main

import (
	"github.com/guru-invest/guru.corporate.actions/src/core"
	"github.com/guru-invest/guru.corporate.actions/src/crossCutting/options"
)

func init() {
	options.OPTIONS.Load()
}

func main() {
	core.Run()
}
