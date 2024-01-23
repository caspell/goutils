package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

var sequence int

func init() {
	sequence = 0
}

func GenerateHash(input string, limit int) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])[:limit]
}

func GetHashBasedTime() string {
	sequence++
	val := time.Now().UnixMilli()
	str := fmt.Sprintf("%s%04d", strconv.FormatInt(val, 10), sequence%10000)
	log.Println(str)
	return strings.ToUpper(GenerateHash(str, 16))
}

func GetSerialNumber() string {
	hashValue := GetHashBasedTime()
	result := fmt.Sprintf("%s-%s-%s-%s", hashValue[:4],
		hashValue[4:8],
		hashValue[8:12],
		hashValue[12:16])
	return result
}
