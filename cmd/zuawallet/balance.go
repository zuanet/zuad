package main

import (
	"context"
	"fmt"

	"github.com/zuanet/zuad/cmd/zuawallet/daemon/client"
	"github.com/zuanet/zuad/cmd/zuawallet/daemon/pb"
	"github.com/zuanet/zuad/cmd/zuawallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatZua(addressBalance.Available), utils.FormatZua(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, ZUA %s %s%s\n", utils.FormatZua(response.Available), utils.FormatZua(response.Pending), pendingSuffix)

	return nil
}
