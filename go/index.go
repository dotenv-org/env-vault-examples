package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

const nonceBytes = 12

func main() {
	key := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	message := "HELLO"

	cipherText, err := encrypt(key, message)
	if err != nil {
		log.Fatal("error in encrypt: ", err)
	}
	fmt.Println("Cipher text:", cipherText)

	plainTextBytes, err := decrypt(key, cipherText)
	if err != nil {
		log.Fatal("error in decrypt: ", err)
	}
	plainText := string(plainTextBytes)
	fmt.Println("Plain text:", plainText)

	if plainText != message {
		fmt.Println("ERROR: plaintext doesn't match original message")
	}
}

// implement
func encrypt(key string, plainText string) (string, error) {
	if len(key) < 64 {
		return "", errors.New("INVALID_DOTENV_KEY: Key part must be 64 characters long (or more)")
	}

	// Set up key.
	keyBytes, _ := hex.DecodeString(key)

	// Generate random nonce.
	nonceBytes := generateNonce()

	// Set up cipher.
	block, _ := aes.NewCipher(keyBytes)
	aesgcm, _ := cipher.NewGCM(block)

	// Encrypt.
	cipherTextBytes := aesgcm.Seal(nil, nonceBytes, []byte(plainText), nil)

	// Prepend nonce and base64 encode.
	messageBytes := append(nonceBytes, cipherTextBytes...)
	return base64.StdEncoding.EncodeToString(messageBytes), nil
}

// Decrypt a single encrypted environment string using the supplied
// key. The cipher is AES-GCM, and the first 12 bytes of the
// ciphertext are used as the nonce value.
func decrypt(key string, cipherText string) ([]byte, error) {
	if len(key) < 64 {
		return nil, errors.New("INVALID_DOTENV_KEY: Key part must be 64 characters long (or more)")
	}

	// Set up key.
	keyBytes, _ := hex.DecodeString(key)

	// Base64 decode input.
	cipherTextBytes, _ := base64.StdEncoding.DecodeString(cipherText)

	// Extract nonce.
	nonceBytes := cipherTextBytes[:12]

	// Extract cipher text.
	cipherTextBytes = cipherTextBytes[12:]

	// Set up cipher.
	block, _ := aes.NewCipher(keyBytes)
	aesgcm, _ := cipher.NewGCM(block)

	// Decrypt.
	plainText, _ := aesgcm.Open(nil, nonceBytes, cipherTextBytes, nil)

	return plainText, nil
}

func generateNonce() []byte {
	nonce := make([]byte, nonceBytes)
	_, _ = rand.Read(nonce)
	return nonce
}
