package base

import (
	"log"

	"github.com/prateekgupta3991/blockchain-experiment/util"
)

type NodeOpn interface {
	FindSuccessor()
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