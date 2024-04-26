package ready

import (
	"github.com/zuanet/zuad/infrastructure/logger"
	"github.com/zuanet/zuad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.ZuaoutineWrapperFunc(log)
