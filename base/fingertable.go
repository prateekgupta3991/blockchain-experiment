package base

import (
	"log"

	"github.com/prateekgupta3991/blockchain-experiment/util"
)

type TableOperation interface {
	InsertValue(k interface{}) error
	ReadValue(k interface{}) (interface{}, error)
}

type Table struct {
	Entry map[int64]TableVal
	KV map[int64]int64
}

type TableVal struct {
	Nid int64
	StartVal int64
	EndVal int64
	Succesor int64
}

func NewFingertable() *Table {
	return &Table{
		Entry: make(map[int64]TableVal, 10),
		KV:  make(map[int64]int64, 10),
	}
}

func (t *Table) InsertValue(k interface{}) error {
	if hash, err := util.CalcHash(k); err != nil {
		log.Printf("Error : %s", err.Error())
		return err
	} else {
		if nid, err := util.GetNid(); err != nil {
			log.Printf("Error : %s", err.Error())
			return err
		} else {
			// t.Entry[hash] = TableVal{
			// 	Nid: nid,
			// 	StartVal: v,
			// 	EndVal: ,
			// }
			log.Printf("hash - %v, nid - %v", hash, nid)
			val := int64(k.(int))
			t.KV[val] = val
			return nil
		}
	}
}

func (t *Table) ReadValue(k interface{}) (interface{}, error) {
	if hash, err := util.CalcHash(k); err != nil {
		log.Printf("Error : %s", err.Error())
		return nil, err
	} else {
		log.Printf("hash - %v", hash)
		val := k.(int64)
		return t.KV[val], nil
	}
}