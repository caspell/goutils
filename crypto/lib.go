package crypto

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"log"

	"golang.org/x/crypto/pbkdf2"
)

const _SALT_SIZE = 64
const _HASH_ITERATION = 10000
const _HASH_KEY_LENGTH = 50

func CreateRandomSalt() string {
	var salt = make([]byte, _SALT_SIZE)

	_, err := rand.Read(salt[:])
	if err != nil {
		log.Println(err.Error())
	}

	return base64.StdEncoding.EncodeToString(salt)
}

func HashPassword(password string, saltValue string) string {
	salt, err := base64.StdEncoding.DecodeString(saltValue)
	if err != nil {
		log.Println("Failed to base64 decoding.", err)
		return ""
	}

	hash := pbkdf2.Key([]byte(password), salt, _HASH_ITERATION, _HASH_KEY_LENGTH, sha512.New)

	return base64.StdEncoding.EncodeToString(hash)
}

func ComparePassword(hashedPassword, currPassword string, salt string) bool {
	var currPasswordHash = HashPassword(currPassword, salt)
	return hashedPassword == currPasswordHash
}

func Base64Encode(val string) string {
	return base64.StdEncoding.EncodeToString([]byte(val))
}
