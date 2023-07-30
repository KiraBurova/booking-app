package auth

import (
	"encoding/json"
	"log"
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
		log.Panic(err)
	}

	compareErr := comparePaswords(data.Password, user.Password)

	if compareErr != nil {
		log.Panic(compareErr)
		w.WriteHeader(http.StatusUnauthorized)
	}

	sessionToken := uuid.NewString()

	createErr := authRepo.Create(sessionToken)

	if createErr != nil {
		log.Panic(createErr)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: sessionToken,
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authRepo := NewRepository(db.DbInstance)

	cookie, err := r.Cookie("session_token")

	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	deleteErr := authRepo.Delete(cookie.Value)

	if deleteErr != nil {
		log.Panic(deleteErr)
		w.WriteHeader(http.StatusBadRequest)
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: "",
	})
}
