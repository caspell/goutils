package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"log"
	"time"
)

func init() {
	log.Println(time.Now())
}

func RSA() {

	// salt := CreateRandomSalt()
	// log.Println(salt)
	// pwd := HashPassword("test123", salt)
	// log.Println(pwd)

	priv, _ := rsa.GenerateKey(rand.Reader, 1024)

	privData1 := x509.MarshalPKCS1PrivateKey(priv)
	privData8, _ := x509.MarshalPKCS8PrivateKey(priv)

	priv1, _ := x509.ParsePKCS1PrivateKey(privData1)
	parsedData, _ := x509.ParsePKCS8PrivateKey(privData8)

	priv8 := parsedData.(*rsa.PrivateKey)

	value := []byte("test ok")

	cip, _ := rsa.EncryptPKCS1v15(rand.Reader, &priv1.PublicKey, value)

	plain, _ := rsa.DecryptPKCS1v15(rand.Reader, priv8, cip)

	log.Println(fmt.Printf("%x", privData1))
	log.Println(fmt.Printf("%x", privData8))

	log.Println(string(plain))

}
