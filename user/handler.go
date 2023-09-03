package user

import (
	"encoding/json"
	"net/http"
	"timezone-converter/db"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var u User
	json.NewDecoder(r.Body).Decode(&u)

	defaultTimeslots := `{
		"9:00":  {"Booked": false},
		"10:00": {"Booked": false},
		"11:00": {"Booked": false},
		"12:00": {"Booked": false},
		"13:00": {"Booked": false},
		"14:00": {"Booked": false},
		"15:00": {"Booked": false},
		"16:00": {"Booked": false}
	}`
	u.Timeslots = defaultTimeslots
	u.Id = uuid.NewString()
	password, errorSetPassword := u.setPassword()
	u.Password = password

	if errorSetPassword != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	userExists, userExistsError := repo.UserExists(u.Username)

	if userExistsError != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if userExists == true {
		w.WriteHeader(http.StatusConflict)
		return
	}

	err := repo.Create(u)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(u)

}

type BookTimeData struct {
	Id   string `json:"id"`
	Time string `json:"time"`
}

func BookTime(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var data BookTimeData
	json.NewDecoder(r.Body).Decode(&data)

	user, err := repo.GetById(data.Id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	var timeSlots map[string]interface{}
	err = json.Unmarshal([]byte(user.Timeslots), &timeSlots)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if slot, ok := timeSlots[data.Time]; ok {
		var booked = slot.(map[string]interface{})["Booked"]

		if booked == true {
			// or 409 Conflict ?
			w.WriteHeader(http.StatusForbidden)
		}

		if booked == false {
			slot.(map[string]interface{})["Booked"] = true
		}

		marshaledJson, _ := json.Marshal(timeSlots)
		user.Timeslots = string(marshaledJson)
		repo.Update(user)
	}

}
