package main

import (
	"context"
	"fmt"
	"github.com/zuanet/zuad/cmd/zuawallet/daemon/client"
	"github.com/zuanet/zuad/cmd/zuawallet/daemon/pb"
)

func newAddress(conf *newAddressConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()

	response, err := daemonClient.NewAddress(ctx, &pb.NewAddressRequest{})
	if err != nil {
		return err
	}
//my-add New address:\n%s\n
	fmt.Printf("%s", response.Address)
	return nil
}
