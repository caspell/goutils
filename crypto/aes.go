package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"strings"
)

type Crypto interface {
	Encrypt(plainText string) (string, error)
	Decrypt(cipherIvKey string) (string, error)
}

type _crypto struct {
	cipherKey   string
	cipherIvKey string
}

func (c _crypto) Encrypt(plainText string) (string, error) {
	if strings.TrimSpace(plainText) == "" {
		return plainText, nil
	}
	block, err := aes.NewCipher([]byte(c.cipherKey))
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCBCEncrypter(block, []byte(c.cipherIvKey))
	paddedPlainText := padPKCS7([]byte(plainText), encrypter.BlockSize())
	cipherText := make([]byte, len(paddedPlainText))
	encrypter.CryptBlocks(cipherText, paddedPlainText)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func (c _crypto) Decrypt(cipherText string) (string, error) {
	if strings.TrimSpace(cipherText) == "" {
		return cipherText, nil
	}
	decodedCipherText, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(c.cipherKey))
	if err != nil {
		return "", err
	}
	decrypter := cipher.NewCBCDecrypter(block, []byte(c.cipherIvKey))
	plainText := make([]byte, len(decodedCipherText))
	decrypter.CryptBlocks(plainText, decodedCipherText)
	return string(trimPKCS5(plainText)), nil
}

func NewNiceCrypto(cipherKey, cipherIvKey string) (Crypto, error) {
	if keyLength := len(cipherKey); keyLength != 32 {
		return nil, aes.KeySizeError(keyLength)
	}
	if ivKeyLength := len(cipherIvKey); ivKeyLength != 16 {
		return nil, aes.KeySizeError(ivKeyLength)
	}
	return &_crypto{cipherKey, cipherIvKey}, nil
}

func padPKCS7(plainText []byte, blockSize int) []byte {
	padding := blockSize - len(plainText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plainText, padText...)
}

func trimPKCS5(text []byte) []byte {
	padding := text[len(text)-1]
	return text[:len(text)-int(padding)]
}
