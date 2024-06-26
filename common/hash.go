package common

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hash struct{}

func NewMd5Hash() *md5Hash {
	return &md5Hash{}
}

func (h *md5Hash) Hash(data string) string {
	hashery := md5.New()
	hashery.Write([]byte(data))

	return hex.EncodeToString(hashery.Sum(nil))
}
