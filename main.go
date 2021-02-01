package main

import (
	"fmt"
	"log"
	"net"

	"github.com/prateekgupta3991/blockchain-experiment/base"
	genSer "github.com/prateekgupta3991/vault/rpc/gen"
	svr "github.com/prateekgupta3991/blockchain-experiment/servers"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is P2P stuff")
	pnode, _ := base.NewNode()
	log.Println("Node id : ", pnode.Id)
	pnode.Ft.InsertValue(123)
	pnode.Ft.InsertValue(134)
	pnode.Ft.InsertValue(123)

	address := "0.0.0.0:8097"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Printf("Server is listening on %v ...", address)

	s := grpc.NewServer()
	pSrvr := &svr.Srvr{
		Primary: pnode,
	}
	genSer.RegisterNodeOpsServer(s, pSrvr)
	s.Serve(lis)
}
