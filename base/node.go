package base

import (
	"log"

	"github.com/prateekgupta3991/blockchain-experiment/util"
)

type NodeOpn interface {
	FindSuccessor() *Node
	InsertNode(ip string) *Node
}

type Node struct {
	Id string
	Ft *Table
}

func NewNode() (*Node, error) {
	if idn, err := util.GetNid(); err != nil {
		log.Printf("Error : %s", err.Error())
		return nil, err
	} else {
		return &Node{
			Id: idn,
			Ft: NewFingertable(),
		}, nil
	}
}

func (n *Node) FindSuccessor() *Node {
	return nil
}

func (n *Node) InsertNode(ip string) *Node {
	return nil
}