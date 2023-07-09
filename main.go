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
	myRouter.HandleFunc("/book_time/{userId}", bookTime).Methods("POST")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Timezones = []string{"Europe/Berlin", "Africa/Abidjan", "Africa/Addis_Ababa"}

	createDb()
	handleRequests()
}

type TimeslotStatus struct {
	Booked bool `json:"booked"`
}

/* USER */
type User struct {
	Id        string                    `json:"id"`
	Username  string                    `json:"username"`
	Password  string                    `json:"password"`
	Timeslots map[string]TimeslotStatus `json:"timeslots"`
}

const create string = `create table users(id integer primary key autoincrement, username text, password text, timeslots blob);`

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
	defaultTimeslots := map[string]TimeslotStatus{
		"9:00":  {Booked: false},
		"10:00": {Booked: false},
		"11:00": {Booked: false},
		"12:00": {Booked: false},
		"13:00": {Booked: false},
		"14:00": {Booked: false},
		"15:00": {Booked: false},
		"16:00": {Booked: false},
	}

	u.Timeslots = defaultTimeslots
	jsonWithTimeslots, _ := json.Marshal(defaultTimeslots)
	json.NewDecoder(r.Body).Decode(&u)
	db, _ := sql.Open("sqlite3", "./users.db")
	stmt, _ := db.Prepare("INSERT INTO users(username, password, timeslots) values(?,?,?)")
	res, _ := stmt.Exec(u.Username, u.Password, jsonWithTimeslots)
	id, _ := res.LastInsertId()

	fmt.Println(id)
	json.NewEncoder(w).Encode(u)
}

func bookTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	db, _ := sql.Open("sqlite3", "./users.db")
	row := db.QueryRow("SELECT * FROM users WHERE id=?", userId)
	user := User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Timeslots)

	// TODO: Timeslots seems to be empty there, figure out why
	fmt.Println(user.Timeslots)
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
