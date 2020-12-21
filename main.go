package main

import (
	"fmt"
	"log"

	"github.com/prateekgupta3991/blockchain-experiment/base"
)

func main() {
	fmt.Println("This is P2P stuff")
	node, _ := base.NewNode()
	log.Println("Node id : ", node.Id)
	node.Ft.InsertValue("123", "qwerty")
	node.Ft.InsertValue("134", "hey")
	node.Ft.InsertValue("123", "qwerty")
}