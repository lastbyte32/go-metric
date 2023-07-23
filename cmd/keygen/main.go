// package main Генерация пары ключей.
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"
)

const (
	privateKeyFilePath = "./private_key.pem"
	publicKeyFilePath  = "./public_key.pem"
)

func main() {
	log.Println("generating private key and public key...")

	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		log.Fatal(err)
	}

	pemPrivateFile, err := os.Create(privateKeyFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer pemPrivateFile.Close()

	if err := pem.Encode(pemPrivateFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}); err != nil {
		log.Fatal(err)
	}

	pemPublicFile, err := os.Create(publicKeyFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer pemPublicFile.Close()
	if err := pem.Encode(pemPublicFile, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	}); err != nil {
		log.Fatal(err)
	}
	log.Println("generated keys successfully")
	log.Printf("public key: [%s] and private key: [%s]\n", publicKeyFilePath, privateKeyFilePath)
}
