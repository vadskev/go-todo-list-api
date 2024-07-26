package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func CreateHash(key string) string {
	h := hmac.New(sha256.New, []byte(os.Getenv("TODO_PASSWORD")))
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}
