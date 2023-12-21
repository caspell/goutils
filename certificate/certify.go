package certificate

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func generateRSAKeys() error {
	// Generate RSA key pair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("failed to generate private key: %s", err)
	}

	// Extract public key from private key
	publicKey := &privateKey.PublicKey

	// Encode private key to PEM format
	privateKeyFile, err := os.Create("private_key.pem")
	if err != nil {
		return fmt.Errorf("failed to create private key file: %s", err)
	}
	defer privateKeyFile.Close()

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return fmt.Errorf("failed to write private key to file: %s", err)
	}

	// Encode public key to PEM format
	publicKeyFile, err := os.Create("public_key.pem")
	if err != nil {
		return fmt.Errorf("failed to create public key file: %s", err)
	}
	defer publicKeyFile.Close()

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %s", err)
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	if err := pem.Encode(publicKeyFile, publicKeyPEM); err != nil {
		return fmt.Errorf("failed to write public key to file: %s", err)
	}

	return nil
}

func Main() {
	if err := generateRSAKeys(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("RSA key pair generated successfully")
}
