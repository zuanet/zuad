package main

import (
	"github.com/zuanet/zuad/infrastructure/logger"
	"github.com/zuanet/zuad/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("JSTT")
	spawn      = panics.GoroutineWrapperFunc(log)
)
