package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	rpc "github.com/prateekgupta3991/vault/rpc/gen"
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
	msg := rpc.KeyDetails{
		Hash: "123",
	}
	if resp, err := c.FindDestNode(context.Background(), &msg); err != nil {
		log.Fatalf("Error FindDestNode : %s", err)
	} else {
		log.Printf("Response from server: %s", resp.Nid)
	}

	nd := rpc.NodeDetails{
		Ip: "123.123.123.123",
		M:  3,
	}
	if resp, err := c.InsertNode(context.Background(), &nd); err != nil {
		log.Fatalf("Error InsertNode : %s", err)
	} else {
		log.Printf("Response from server: %s - %s", resp.Ip, resp.Nid)
	}

	bulkNodeInserts(c)
}

func bulkNodeInserts(c rpc.NodeOpsClient) {
	stream, err := c.InsertNodes(context.Background())
	if err != nil {
		fmt.Errorf("Error encountered in client : %v", err)
		return
	}

	waitChan := make(chan struct{})

	go func() {
		for i := 1; i < 50000; i++ {
			stream.Send(&rpc.NodeDetails{
				Ip: strconv.Itoa(i),
				M:  int64(i),
			})
		}
		fmt.Println("Request sent Timestamp : ", time.Now())
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("Error encountered while receiving response in client : ", err)
				return
			}

			fmt.Println("Server response : ", res)
		}
		fmt.Println("Response received Timestamp : ", time.Now())
		close(waitChan)
	}()
	<-waitChan
}
