package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt 散列加密
func Encrypt(plaintext string) (string, error) {
	ciphertext, err := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(ciphertext), nil
}

// ComparePassword 比较密码
func ComparePassword(a, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}
