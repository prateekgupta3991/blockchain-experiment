package main

import (
	"context"
	"log"

	rpc "github.com/prateekgupta3991/blockchain-experiment/rpc/gen"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	var cerr error
	if conn, cerr = grpc.Dial("0.0.0.0:8097", grpc.WithInsecure()); cerr != nil {
		log.Fatalf("did not connect: %s", cerr)
	}
	defer conn.Close()

	c := rpc.NewNodeOpsClient(conn)
	msg := rpc.KeyDetails {
		Hash: "123",
	}
	if resp, err := c.FindDestNode(context.Background(), &msg); err != nil {
		log.Fatalf("Error FindDestNode : %s", err)
	} else {
		log.Printf("Response from server: %s", resp.Nid)
	}

	nd := rpc.NodeDetails{
		Ip: "123.123.123.123",
		M: 3,
	}
	if resp, err := c.InsertNode(context.Background(), &nd); err != nil {
		log.Fatalf("Error InsertNode : %s", err)
	} else {
		log.Printf("Response from server: %s - %s", resp.Ip, resp.Nid)
	}
}
