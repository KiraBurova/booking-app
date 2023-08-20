package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func comparePaswords(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err
}

func getCookie(r *http.Request) (*http.Cookie, error) {
	return r.Cookie("session_token")
}
