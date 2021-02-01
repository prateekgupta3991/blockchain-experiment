package servers

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/prateekgupta3991/blockchain-experiment/base"
	rpc "github.com/prateekgupta3991/vault/rpc/gen"
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

func (s *Srvr) InsertNodes(stream rpc.NodeOps_InsertNodesServer) error {
	fmt.Println("InsertNodes Function")
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Errorf("Error encountered : %v", err)
			return err
		}

		fmt.Println("Got a request from ip : ", req.Ip)

		if res := stream.Send(&rpc.NewNodeDetails{
			Nid: strconv.Itoa(int(req.M)),
			Ip: req.Ip,
		}); res != nil {
			fmt.Println("Error when response was sent to the client: ", res)
		}
	}
	return nil
}
