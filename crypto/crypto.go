package crypto

import (
	"log"
	"time"
)

func init() {
	log.Println(time.Now())
}

func Main() {
	cipherKey := "_TTM__CIPHER_KEY_0123456789012__"
	cipherIvKey := "__CIPHERIVKEY0__"

	log.Println(len(cipherKey), len(cipherIvKey))

	c, err := NewNiceCrypto(cipherKey, cipherIvKey)

	if err != nil {
		panic(err)
	}

	encrypted, _ := c.Encrypt("test ok")

	log.Println(encrypted)

	decrypted, _ := c.Decrypt(encrypted)

	log.Println(decrypted)

}
