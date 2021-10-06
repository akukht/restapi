package hashing

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//MakePass generate password
func MakePass(password string) (string, error) {
	if password != "" {
		bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		return string(bytes), err
	}

	return "", errors.New("empty password")
}

//CheckPasswordHash compare passwords
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
