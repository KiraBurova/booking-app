package user

import (
	"database/sql"
	"encoding/json"
	"errors"
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

	userExists, err := repo.userExists(u.Username)

	if err != nil && errors.Is(err, sql.ErrNoRows) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userExists {
		w.WriteHeader(http.StatusConflict)
		return
	} else {
		err := repo.Create(u)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(u)
	}

}
