package main

import (
	"log"

	"gitlab.antline.com/golib/crypto.git"
)

func main() {

	key := ""
	akey := ""

	value := ""

	aes := crypto.NewAES256GSM(key, akey)
	if enc, err := aes.Encrypt(value); err != nil {
		log.Fatal(err)
	} else {
		log.Println("result: ", len(enc))
	}

}
