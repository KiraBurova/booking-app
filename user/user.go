package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}


func (u *User) setPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	return string(bytes), err
}
