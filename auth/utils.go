package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func comparePaswords(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}
