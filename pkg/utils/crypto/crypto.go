package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

type cryptor struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewEncryptor(fileName string) (*cryptor, error) {
	if fileName == "nil" {
		return nil, fmt.Errorf("public_key.pem not found")
	}
	c := &cryptor{}
	file, err := c.readPem(fileName)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS1PublicKey(file)
	if err != nil {
		return nil, err
	}

	return &cryptor{
		publicKey: key,
	}, nil
}

func NewDecryptor(fileName string) (*cryptor, error) {
	if fileName == "nil" {
		return nil, fmt.Errorf("private_key.pem not found")
	}
	c := &cryptor{}
	file, err := c.readPem(fileName)
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(file)
	if err != nil {
		return nil, err
	}

	return &cryptor{
		privateKey: key,
	}, nil
}

func (c *cryptor) Encrypt(message []byte) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, c.publicKey, message, nil)
	if err != nil {
		return nil, err
	}
	return encryptedBytes, nil
}

func (c *cryptor) Decrypt(encryptedMessage []byte) ([]byte, error) {
	return c.privateKey.Decrypt(nil, encryptedMessage, &rsa.OAEPOptions{Hash: crypto.SHA256})
}

func (c *cryptor) readPem(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(file)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block from %s", fileName)
	}
	return block.Bytes, err
}
