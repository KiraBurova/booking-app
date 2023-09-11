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

	router.HandleFunc("/api/register", user.Register).Methods("POST")
	router.HandleFunc("/api/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/logout", auth.Logout).Methods("POST")

	api := router.PathPrefix("/api").Subrouter()
	api.Use(middlewares()...)

	api.Path("/timezones").Handler(http.HandlerFunc(timezone.ListTimezones))
	api.Path("/timezone").Handler(http.HandlerFunc(timezone.AddTimezone))
	api.Path("/convert_timezone").Handler(http.HandlerFunc(timezone.ConvertTimezone))
	api.Path("/book_timeslot").Handler(http.HandlerFunc(user.BookTimeslot))

	err := http.ListenAndServe(":10000", router)

	if err != nil {
		log.Fatal(err)
	}

}

func middlewares() []mux.MiddlewareFunc {
	return []mux.MiddlewareFunc{
		auth.ValidationMiddleware,
	}
}

func main() {
	db.InitDb("users.db")
	handleRequests()
}
