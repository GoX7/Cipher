package cipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

func cashKey(key string) []byte {
	hKay := sha256.Sum256([]byte(key))
	return hKay[:]
}

func Encrypt(key string, data []byte) (string, error) {
	block, err := aes.NewCipher(cashKey(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherData := gcm.Seal(nonce, nonce, data, nil)
	return base64.URLEncoding.EncodeToString(cipherData), nil
}

func Decrypt(key string, cipherText string) (string, error) {
	cipherData, err := base64.URLEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(cashKey(key))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce, text := cipherData[:gcm.NonceSize()], cipherData[gcm.NonceSize():]
	cipherOpen, err := gcm.Open(nil, nonce, text, nil)
	if err != nil {
		return "", err
	}

	return string(cipherOpen), nil
}
