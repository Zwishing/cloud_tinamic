package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
)

const DefaultSalt = "66e38420-2cd8-8001-91cf-501e513d3bbc"

func RandomSalt() string {
	saltBytes := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, saltBytes)
	if err != nil {
		return DefaultSalt
	}
	salt := base64.StdEncoding.EncodeToString(saltBytes)
	return salt
}

func CreateHashPassword(password string, salt string) string {
	// 加盐处理
	toHash := password + salt
	h := sha256.New()
	h.Write([]byte(toHash))
	hashPassword := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return hashPassword
}

// ValidatePassword 验证密码是否正确
func ValidatePassword(password, salt, hashPassword string) bool {
	computedHash := CreateHashPassword(password, salt)
	return hashPassword == computedHash
}
