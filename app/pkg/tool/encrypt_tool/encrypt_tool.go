package encrypt_tool

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// EncryptString encrypts a string using AES-GCM and returns a base64 encoded ciphertext.
func Encrypt(encryptData []byte, key string) (string, error) {
	if len(key) != 32 {
		key = adjustKeyLength(key, 32)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, encryptData, nil)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// DecryptString decrypts a base64 encoded ciphertext string using AES-GCM and returns the plaintext.
func Decrypt(encryptedText, key string) ([]byte, error) {
	if len(key) != 32 {
		key = adjustKeyLength(key, 32)
	}

	// Decode the base64 encoded ciphertext
	cipherText, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	// Extract the nonce and actual ciphertext
	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]

	// Decrypt the ciphertext
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func adjustKeyLength(key string, keyLength int) string {
	keyBytes := []byte(key)
	if len(keyBytes) < keyLength {
		paddedKey := make([]byte, keyLength)
		copy(paddedKey, keyBytes)
		return string(paddedKey)
	}
	return string(keyBytes[:keyLength])
}
