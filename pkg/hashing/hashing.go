package hashing

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func MakePass(password string) (string, error) {
	if password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err
	}

	return "", errors.New("Empty password")
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
