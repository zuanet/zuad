package main

import (
	"fmt"
	"os"

	"github.com/zuanet/zuad/infrastructure/logger"
	"github.com/zuanet/zuad/stability-tests/common"
	"github.com/zuanet/zuad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("KSSA")
	spawn      = panics.ZuaoutineWrapperFunc(log)
)

func initLog(logFile, errLogFile string) {
	level := logger.LevelDebug
	if activeConfig().LogLevel != "" {
		var ok bool
		level, ok = logger.LevelFromString(activeConfig().LogLevel)
		if !ok {
			fmt.Fprintf(os.Stderr, "Log level %s doesn't exists", activeConfig().LogLevel)
			os.Exit(1)
		}
	}
	log.SetLevel(level)
	common.InitBackend(backendLog, logFile, errLogFile)
}
