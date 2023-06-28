package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

var Timezones []string

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/timezones", listTimezones)
	myRouter.HandleFunc("/timezone", addTimezone).Methods("POST")
	myRouter.HandleFunc("/convert_timezone", convertTimezone).Methods("POST")

	myRouter.HandleFunc("/register", register).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Timezones = []string{"Europe/Berlin", "Africa/Abidjan", "Africa/Addis_Ababa"}

	createDb()
	handleRequests()
}

/* USER */
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const create string = `create table users(username text, password text);`

func createDb() error {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return err
	}
	if _, err := db.Exec(create); err != nil {
		return err
	}
	return nil
}

func register(w http.ResponseWriter, r *http.Request) {
	// get posted data for user and store into db
	var u User

	json.NewDecoder(r.Body).Decode(&u)

	db, _ := sql.Open("sqlite3", "./users.db")

	stmt, _ := db.Prepare("INSERT INTO users(username, password) values(?,?)")

	res, _ := stmt.Exec(u.Username, u.Password)

	id, _ := res.LastInsertId()

	fmt.Println(id)

	json.NewEncoder(w).Encode(u)
}

/* TIMEZONES */
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
func convertTimezone(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	timezone, _ := ioutil.ReadAll(r.Body)

	loc, _ := time.LoadLocation(string(timezone))

	fmt.Printf("Time in %v: %s\n", string(timezone), now.In(loc))
}
