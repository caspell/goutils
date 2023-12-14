package main

import (
	"log"

	"gitlab.antline.com/golib/crypto.git"
)

func main() {

	key := "_TTM__CIPHER_KEY_0123456789012__"
	akey := "caspell@naver.com"

	value := "test12341!"

	aes := crypto.NewAES256GSM(key, akey)
	if enc, err := aes.Encrypt(value); err != nil {
		log.Fatal(err)
	} else {
		log.Println("result: ", len(enc))
	}

}
