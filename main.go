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

func main() {
	fmt.Println("This is P2P stuff")
	node, _ := base.NewNode()
	log.Println("Node id : ", node.Id)
	node.Ft.InsertValue("123", "qwerty")
	node.Ft.InsertValue("134", "hey")
	node.Ft.InsertValue("123", "qwerty")

	address := "0.0.0.0:8097"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	genSer.RegisterDestinationNodeServer(s, &svr.Srvr{})
	s.Serve(lis)
}
