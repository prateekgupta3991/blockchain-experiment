package util

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func CalcHash(k interface{}) (string, error) {
	val := md5.New().Sum([]byte(strconv.Itoa(k.(int))))
	return hex.EncodeToString(val[:]), nil
}
