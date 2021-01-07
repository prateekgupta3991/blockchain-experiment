package main

import (
	"fmt"
	"log"
	"net"

	"github.com/prateekgupta3991/blockchain-experiment/base"
	genSer "github.com/prateekgupta3991/blockchain-experiment/rpc/gen"
	svr "github.com/prateekgupta3991/blockchain-experiment/servers"
	"google.golang.org/grpc"
)

var Pnode *base.Node

func main() {
	fmt.Println("This is P2P stuff")
	Pnode, _ = base.NewNode()
	log.Println("Node id : ", Pnode.Id)
	Pnode.Ft.InsertValue(123)
	Pnode.Ft.InsertValue(134)
	Pnode.Ft.InsertValue(123)

	address := "0.0.0.0:8097"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	genSer.RegisterNodeOpsServer(s, &svr.Srvr{})
	s.Serve(lis)
}
