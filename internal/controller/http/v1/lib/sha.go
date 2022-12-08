package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type SHAHasher struct {
	Hasher    hash.Hash
	SecretKey []byte
}

func NewSHA() hash.Hash {
	return sha256.New()
}

func (sha SHAHasher) HashPassword(password string) string {
	sha.Hasher.Write([]byte(password))
	return hex.EncodeToString(sha.Hasher.Sum(sha.SecretKey))
}
