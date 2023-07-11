package main

import (
	"log"
	"net/http"
	"timezone-converter/db"
	"timezone-converter/timezone"
	"timezone-converter/user"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/timezones", timezone.ListTimezones)
	myRouter.HandleFunc("/timezone", timezone.AddTimezone).Methods("POST")
	myRouter.HandleFunc("/convert_timezone", timezone.ConvertTimezone).Methods("POST")
	myRouter.HandleFunc("/register", user.Register).Methods("POST")
	myRouter.HandleFunc("/book_time/{userId}", user.BookTime).Methods("POST")

	err := http.ListenAndServe(":10000", myRouter)

	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	db.InitDb()
	handleRequests()
}
