package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const SERIAL_NUMBER_LENGTH = 16

var sequence int

func init() {
	sequence = 0
}

func GenerateHash(input string, limit int) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])[:limit]
}

func GetHashBasedTime(limit int) string {
	sequence++
	val := time.Now().UnixMilli()
	str := fmt.Sprintf("%s%04d", strconv.FormatInt(val, 10), sequence%10000)
	return strings.ToUpper(GenerateHash(str, limit))
}

func GetSerialNumber() string {

	hashValue := GetHashBasedTime(SERIAL_NUMBER_LENGTH)
	result := fmt.Sprintf("%s-%s-%s-%s", hashValue[:4],
		hashValue[4:8],
		hashValue[8:12],
		hashValue[12:])
	return result
}
