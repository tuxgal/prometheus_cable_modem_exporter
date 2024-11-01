package main

import (
	"github.com/tuxgal/tuxlog"
	"github.com/tuxgal/tuxlogi"
)

var (
	log = buildLogger()
)

func buildLogger() tuxlogi.Logger {
	config := tuxlog.NewConsoleLoggerConfig()
	config.MaxLevel = tuxlog.LvlDebug
	return tuxlog.NewLogger(config)
}
