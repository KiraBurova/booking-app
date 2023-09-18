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

	json.NewEncoder(w).Encode(u)

}

func BookTimeslot(w http.ResponseWriter, r *http.Request) {

}
