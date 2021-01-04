package main

import (
	"context"
	"log"

	rpc "github.com/prateekgupta3991/blockchain-experiment/rpc/gen"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	var err error
	if conn, err = grpc.Dial("0.0.0.0:8097", grpc.WithInsecure()); err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := rpc.NewDestinationNodeClient(conn)
	var resp *rpc.DestinationNodeDetails
	msg := rpc.KeyDetails {
		Hash: "123",
	}
	if resp, err = c.FindDestNode(context.Background(), &msg); err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", resp.Nid)
}
