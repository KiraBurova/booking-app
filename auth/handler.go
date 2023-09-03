package auth

import (
	"encoding/json"
	"net/http"
	"timezone-converter/db"
	"timezone-converter/user"

	"github.com/google/uuid"
)

func Login(w http.ResponseWriter, r *http.Request) {
	userRepo := user.NewRepository(db.DbInstance)
	authRepo := NewRepository(db.DbInstance)

	var data LoginData
	json.NewDecoder(r.Body).Decode(&data)

	user, err := userRepo.GetByUsername(data.Username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	compareErr := comparePaswords(data.Password, user.Password)

	if compareErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}

	sessionToken := uuid.NewString()

	createErr := authRepo.Create(sessionToken, user.Id)

	if createErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: sessionToken,
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authRepo := NewRepository(db.DbInstance)

	cookie, err := getCookie(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	deleteErr := authRepo.Delete(cookie.Value)

	if deleteErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: "",
	})
}
