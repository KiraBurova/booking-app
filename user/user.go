package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"timezone-converter/db"

	"github.com/gorilla/mux"
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

func Register(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var u User
	defaultTimeslots := `{
		"9:00":  {Booked: false},
		"10:00": {Booked: false},
		"11:00": {Booked: false},
		"12:00": {Booked: false},
		"13:00": {Booked: false},
		"14:00": {Booked: false},
		"15:00": {Booked: false},
		"16:00": {Booked: false},
	}`
	u.Timeslots = defaultTimeslots

	json.NewDecoder(r.Body).Decode(&u)

	err := repo.Create(u)

	if err != nil {
		log.Panic(err)
	}

	json.NewEncoder(w).Encode(u)
}

func BookTime(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	params := mux.Vars(r)
	userId := params["userId"]
	user, err := repo.GetById(userId)

	if err != nil {
		log.Panic(err)
	}

	// for now
	fmt.Println(user)

	// TODO: Timeslots seems to be empty there, figure out why
}
