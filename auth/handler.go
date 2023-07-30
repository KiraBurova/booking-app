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
	db.AttachDb(`ATTACH DATABASE 'sessions.db' as 'sessions'`)

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

	authRepo.Create(sessionToken)

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

	authRepo.Delete(cookie.Value)

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: "",
	})
}
