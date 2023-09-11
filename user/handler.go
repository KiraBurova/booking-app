package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"timezone-converter/db"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var u User
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = uuid.NewString()
	password, errorSetPassword := u.setPassword()
	u.Password = password

	if errorSetPassword != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userExists, userExistsError := repo.UserExists(u.Username)

	if userExists {
		w.WriteHeader(http.StatusConflict)
		return
	}

	if userExistsError == sql.ErrNoRows {
		w.WriteHeader(http.StatusInternalServerError)

		err := repo.Create(u)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	defaultTimeslot := Timeslot{CreatorId: u.Id, InvitedUserId: "", Time: "", Booked: false}
	createTimeslotError := repo.CreateTimeslots(defaultTimeslot)

	if createTimeslotError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(u)

}

func BookTimeslot(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var data Timeslot
	json.NewDecoder(r.Body).Decode(&data)

	timeslot, err := repo.isTimeslotBooked(data)

	// if there are no rows => timeslots were not created yet at all
	if err == sql.ErrNoRows {
		repo.CreateTimeslots(data)
		return
	}

	if timeslot.Booked && timeslot.CreatorId != data.CreatorId {
		http.Error(w, "Somebody else already booked this slot", http.StatusConflict)
		return
	}

	if timeslot.Booked && timeslot.CreatorId == data.CreatorId {
		http.Error(w, "You already booked this slot", http.StatusConflict)
	}
}
