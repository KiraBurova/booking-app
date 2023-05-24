package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var Timezones []string

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/timezones", listTimezones)
	myRouter.HandleFunc("/timezone", addTimezone).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Timezones = []string{"Europe/Berlin", "Africa/Abidjan", "Africa/Addis_Ababa"}

	handleRequests()
}

func getCurrentTimezone() string {
	t := time.Now()
	zone, _ := t.Zone()

	return zone
}

func addTimezone(w http.ResponseWriter, r *http.Request) {
	// read posted timezone
	timezone, _ := ioutil.ReadAll(r.Body)

	// set new timezone to timezones array
	Timezones = append(Timezones, string(timezone))

	json.NewEncoder(w).Encode(timezone)
}
func listTimezones(w http.ResponseWriter, r *http.Request) {
	// get list of timezones in json format
	json.NewEncoder(w).Encode(Timezones)

}
func convertTimezone() {
	// get timezone getCurrentTimezone()
	// convert to any timezone from the listTimezones()
}
