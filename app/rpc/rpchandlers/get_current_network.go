package rpchandlers

import (
	"github.com/zuanet/zuad/app/appmessage"
	"github.com/zuanet/zuad/app/rpc/rpccontext"
	"github.com/zuanet/zuad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}