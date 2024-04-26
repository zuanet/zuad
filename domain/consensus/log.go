package consensus

import (
	"github.com/zuanet/zuad/infrastructure/logger"
	"github.com/zuanet/zuad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.ZuaoutineWrapperFunc(log)
