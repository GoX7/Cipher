# AES-256 GCM Encryption/Decryption Utility

A simple Go package for encrypting and decrypting data using AES-256 in GCM mode.

## Features

- AES-256 encryption with GCM mode (authenticated encryption)
- SHA-256 key derivation
- Random nonce generation for each encryption
- Base64 URL-safe encoding of ciphertext

## Usage

```go
package main

import (
    "github.com/GoX7/cipher"
)

func main() {
	key := "your-secret-key-here"
	plaintext := "sensitive data to encrypt"

	// Encrypt
	encrypted, err := cipher.Encrypt(key, []byte(plaintext))
	if err != nil {
		panic(err)
	}

	// Decrypt
	decrypted, err := cipher.Decrypt(key, encrypted)
	if err != nil {
		panic(err)
	}
}
```
