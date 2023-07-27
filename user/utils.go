package user

import "golang.org/x/crypto/bcrypt"

func comparePaswords(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))

	return err
}
