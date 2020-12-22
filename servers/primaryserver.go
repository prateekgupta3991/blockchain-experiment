package servers

import (
	"context"
	"fmt"

	rpc "github.com/prateekgupta3991/blockchain-experiment/rpc/gen"
	"google.golang.org/grpc/status"
)

type Srvr struct {
}

func (s *Srvr) FindDestNode(ctx context.Context, k *rpc.KeyDetails) (*rpc.DestinationNodeDetails, error) {
	fmt.Println("This server works")
	return nil, status.Errorf(500, "method FindDestNode not implemented")
}
