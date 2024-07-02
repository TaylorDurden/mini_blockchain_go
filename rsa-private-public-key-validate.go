package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"log"
)

func GenerateKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, &privateKey.PublicKey, nil
}

func Encrypt(plaintext []byte, publicKey *rsa.PublicKey) ([]byte, error) {
	label := []byte("")
	message := []byte(plaintext)
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, publicKey, message, label)
	if err != nil {
		return nil, err
	}
	return ciphertext, err
}

func Decrypt(ciphertext []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	label := []byte("")
	hash := sha256.New()
	decryptedPlaintext, err := rsa.DecryptOAEP(hash, rand.Reader, privateKey, ciphertext, label)
	if err != nil {
		return nil, err
	}
	return decryptedPlaintext, nil
}

func PrivatePublicKeyValidte() {
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		log.Fatal("GenerateKeyPair failed..")
	}

	message := []byte("taylorli11111")

	// encrypt by public key
	ciphertext, err := Encrypt(message, publicKey)
	if err != nil {
		log.Fatalf("Encrypt failed: %v", err)
	}

	fmt.Printf("Encrypted text with public key: %x\n", ciphertext)

	// decrypt by private key
	decryptedPlaintext, err := Decrypt(ciphertext, privateKey)
	if err != nil {
		log.Fatalf("Decrypt failed: %v", err)
	}

	fmt.Printf("Decrypted with private key: %s\n", decryptedPlaintext)

}
