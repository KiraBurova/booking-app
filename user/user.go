package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Timeslot struct {
	CreatorId     string `json:"CreatorId"`
	InvitedUserId string `json:"InvitedUserId"`
	Time          string `json:"Time"`
	Booked        bool   `json:"booked"`
}

func (u *User) setPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	return string(bytes), err
}
