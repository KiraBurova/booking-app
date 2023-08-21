package main

import (
	"log"
	"net/http"
	"timezone-converter/auth"
	"timezone-converter/db"
	"timezone-converter/timezone"
	"timezone-converter/user"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/register", user.Register).Methods("POST")
	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.HandleFunc("/logout", auth.Logout).Methods("POST")

	router.HandleFunc("/timezones", timezone.ListTimezones)
	router.HandleFunc("/timezone", timezone.AddTimezone).Methods("POST")
	router.HandleFunc("/convert_timezone", timezone.ConvertTimezone).Methods("POST")
	router.Handle("/book_time", auth.SessionId(http.HandlerFunc(user.BookTime)))

	err := http.ListenAndServe(":10000", router)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	db.InitDb("users.db")
	handleRequests()
}
