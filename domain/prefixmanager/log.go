package prefixmanager

import (
	"github.com/zuanet/zuad/infrastructure/logger"
	"github.com/zuanet/zuad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.ZuaoutineWrapperFunc(log)
