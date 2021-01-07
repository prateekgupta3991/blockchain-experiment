package servers

import (
	"context"
	"fmt"

	"github.com/prateekgupta3991/blockchain-experiment/base"
	rpc "github.com/prateekgupta3991/blockchain-experiment/rpc/gen"
)

type Srvr struct {
	rpc.UnimplementedNodeOpsServer
	Primary *base.Node 
}

func (s *Srvr) FindDestNode(ctx context.Context, k *rpc.KeyDetails) (*rpc.DestinationNodeDetails, error) {
	fmt.Println("This server works")
	return &rpc.DestinationNodeDetails{
		Nid: "0",
	}, nil
}

func (s *Srvr) InsertNode(ctx context.Context, nd *rpc.NodeDetails) (*rpc.NewNodeDetails, error) {
	fmt.Println("This server works")
	// Pnode.Insert("123.123.123.123", 3)
	return &rpc.NewNodeDetails{
		Nid: "0",
		Ip: "0",
	}, nil
}
