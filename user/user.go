package user

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Timeslots string `json:"timeslots"`
}

type TimeslotStatus struct {
	Booked bool `json:"booked"`
}

func (u *User) setPassword() string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}
