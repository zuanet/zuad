package connmanager

import (
	"github.com/zuanet/zuad/infrastructure/logger"
	"github.com/zuanet/zuad/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.ZuaoutineWrapperFunc(log)
