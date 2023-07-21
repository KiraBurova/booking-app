package user

import "golang.org/x/crypto/bcrypt"

func HashPassowrd(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePaswords(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))

	return err
}
