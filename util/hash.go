package util

import (
	"crypto/md5"
	"encoding/hex"
)

func CalcHash(k interface{}) (string, error) {
	val := md5.New().Sum([]byte(k.(string)))
	return hex.EncodeToString(val[:]), nil
}