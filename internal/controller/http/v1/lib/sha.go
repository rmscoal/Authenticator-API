package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"os"
)

type SHAHasher struct {
	Hasher    hash.Hash
	SecretKey []byte
}

func HashPassword(password string) string {
	sha := SHAHasher{
		Hasher:    sha256.New(),
		SecretKey: []byte(os.Getenv("SECRET_KEY_AUTHENTICATOR")),
	}

	sha.Hasher.Write([]byte(password))
	return hex.EncodeToString(sha.Hasher.Sum(sha.SecretKey))
}
