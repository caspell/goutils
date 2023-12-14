package hash

import (
	"crypto/md5"
	"encoding/hex"
)

func GenerateHash(input string, limit int) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])[:limit]
}
