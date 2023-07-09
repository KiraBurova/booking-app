package user

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"timezone-converter/db"

	"github.com/gorilla/mux"
)

type User struct {
	Id        string                    `json:"id"`
	Username  string                    `json:"username"`
	Password  string                    `json:"password"`
	Timeslots map[string]TimeslotStatus `json:"timeslots"`
}

type TimeslotStatus struct {
	Booked bool `json:"booked"`
}

func Register(w http.ResponseWriter, r *http.Request) {
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
	args := []any{u.Username, u.Password, jsonWithTimeslots}
	db.Db_instance.Exec("INSERT INTO users(username, password, timeslots) values(?,?,?)", args...)

	json.NewEncoder(w).Encode(u)
}

func BookTime(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["userId"]
	db, _ := sql.Open("sqlite3", "./users.db")
	row := db.QueryRow("SELECT * FROM users WHERE id=?", userId)
	user := User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Timeslots)

	// TODO: Timeslots seems to be empty there, figure out why
}
