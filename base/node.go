package base

import (
	"log"
	"math"

	"github.com/prateekgupta3991/blockchain-experiment/util"
)

type NodeOpn interface {
	FindSuccessor() *Node
	Insert(ip string, m int) *Node
}

type Node struct {
	Id int64
	Ft *Table
	Ip string
}

func NewNode() (*Node, error) {
	if idn, err := util.GetNid(); err != nil {
		log.Printf("Error : %s", err.Error())
		return nil, err
	} else {
		return &Node{
			Ft: NewFingertable(),
			Ip: idn,
		}, nil
	}
}

func (n *Node) FindSuccessor() *Node {
	return nil
}

func (n *Node) Insert(ip string, m int) *Node {
	keyS := n.Id + int64(math.Pow(float64(2), float64(m-1)))
	keyE := n.Id + int64(math.Pow(float64(2), float64(m)))
	suc := n.Ft.Entry[0].Succesor
	// n.Ft.Entry[0].Succesor = keyS
	node, _ := NewNode()
	node.Ft.Entry[keyS] = TableVal{
		Nid: keyS,
		StartVal: keyS,
		EndVal: keyE,
		Succesor: suc,
	}
	return node
}