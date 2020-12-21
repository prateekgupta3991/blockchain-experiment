package base

import (
	"log"

	"github.com/prateekgupta3991/blockchain-experiment/util"
)

type TableOperation interface {
	InsertValue(k interface{}, v interface{}) error
	ReadValue(k interface{}) (interface{}, error)
}

type Table struct {
	Entry map[string]TableVal
}

type TableVal struct {
	Nid string
	Val interface{}
}

func NewFingertable() *Table {
	return &Table{
		Entry: make(map[string]TableVal, 10),
	}
}

func (t *Table) InsertValue(k interface{}, v interface{}) error {
	if hash, err := util.CalcHash(k); err != nil {
		log.Printf("Error : %s", err.Error())
		return err
	} else {
		if nid, err := util.GetNid(); err != nil {
			log.Printf("Error : %s", err.Error())
			return err
		} else {
			t.Entry[hash] = TableVal{
				Nid: nid,
				Val: v,
			}
			return nil
		}
	}
}

func (t *Table) ReadValue(k interface{}) (interface{}, error) {
	if hash, err := util.CalcHash(k); err != nil {
		log.Printf("Error : %s", err.Error())
		return nil, err
	} else {
		return t.Entry[hash], nil
	}
}