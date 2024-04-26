package client

import (
	"context"
	"github.com/zuanet/zuad/cmd/zuawallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/zuanet/zuad/cmd/zuawallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the zuawalletd server, and returns the client instance
func Connect(address string) (pb.ZuawalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("zuawallet daemon is not running, start it with `zuawallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewZuawalletdClient(conn), func() {
		conn.Close()
	}, nil
}
