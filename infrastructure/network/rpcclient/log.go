package rpcclient

import (
	"github.com/zuanet/zuad/infrastructure/logger"
	"github.com/zuanet/zuad/util/panics"
)

var log = logger.RegisterSubSystem("RPCC")
var spawn = panics.ZuaoutineWrapperFunc(log)
