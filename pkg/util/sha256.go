package util

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256Code(data string) string {
	byteCode:= sha256.Sum256([]byte(data))
	return hex.EncodeToString(byteCode[:])
}