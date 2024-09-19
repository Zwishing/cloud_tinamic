package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

const DefaultSalt = "66e38420-2cd8-8001-91cf-501e513d3bbc"

// RandomSalt generates a cryptographically secure random salt
// It returns the DefaultSalt if there's an error generating random bytes
func RandomSalt() string {
	saltBytes := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, saltBytes); err != nil {
		return DefaultSalt
	}
	return base64.StdEncoding.EncodeToString(saltBytes)
}

// CreateHashPassword creates a salted hash of the password
// It uses SHA-256 for hashing and returns the result as a base64 encoded string
func CreateHashPassword(password, salt string) string {
	h := sha256.New()
	h.Write([]byte(password + salt))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// ValidatePassword checks if the provided password matches the stored hash
// It returns true if the password is correct, false otherwise
func ValidatePassword(password, salt, hashPassword string) bool {
	return CreateHashPassword(password, salt) == hashPassword
}
