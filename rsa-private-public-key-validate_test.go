package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyPair(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPair()
	assert.Nil(t, err)
	assert.NotNil(t, privateKey)
	assert.NotNil(t, publicKey)
}

func TestEncryptDecrypt(t *testing.T) {
	privateKey, publicKey, err := GenerateKeyPair()
	if err != nil {
		panic(err)
	}

	message := []byte("Taylor111111")
	cipherText, err := Encrypt(message, publicKey)
	if err != nil {
		panic(err)
	}
	decryptText, err := Decrypt(cipherText, privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("message: %s\n", message)
	fmt.Printf("decryptText: %s\n", decryptText)
	assert.True(t, reflect.DeepEqual(message, decryptText))
	assert.Equal(t, string(message), string(decryptText))
}
