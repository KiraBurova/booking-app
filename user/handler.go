package user

import (
	"encoding/json"
	"net/http"
	"timezone-converter/db"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var u User
	json.NewDecoder(r.Body).Decode(&u)

	defaultTimeslots := `{
		"9:00":  {"Booked": false},
		"10:00": {"Booked": false},
		"11:00": {"Booked": false},
		"12:00": {"Booked": false},
		"13:00": {"Booked": false},
		"14:00": {"Booked": false},
		"15:00": {"Booked": false},
		"16:00": {"Booked": false}
	}`
	u.Timeslots = defaultTimeslots
	u.Id = uuid.NewString()
	password, errorSetPassword := u.setPassword()
	u.Password = password

	if errorSetPassword != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	userExists, userExistsError := repo.UserExists(u.Username)

	if userExistsError != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if userExists == true {
		w.WriteHeader(http.StatusConflict)
		return
	}

	err := repo.Create(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(u)

}

func BookTime(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	params := mux.Vars(r)
	userId := params["userId"]

	user, err := repo.GetById(userId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var timeSlots map[string]interface{}
	err = json.Unmarshal([]byte(user.Timeslots), &timeSlots)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
