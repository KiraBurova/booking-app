package timeslots

import (
	"encoding/json"
	"net/http"
	"timezone-converter/db"
)

type TimeslotData struct {
	UserId string `json:"userId"`
}

func CreateTimeslots(w http.ResponseWriter, r *http.Request) {
	var timeslotsData TimeslotData
	json.NewDecoder(r.Body).Decode(&timeslotsData)

	repo := NewRepository(db.DbInstance)

	time := [6]string{"10:00", "11:00", "12:00", "13:00", "14:00", "15:00"}

	for i := 0; i < len(time); i++ {
		t := Timeslot{UserId: timeslotsData.UserId, Time: time[i], Booked: false}
		repo.createTimeslots(t)
	}
}

func BookTimeslot(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var data Timeslot
	json.NewDecoder(r.Body).Decode(&data)

	ts, err := repo.getTimeslot(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if ts.Booked {
		w.WriteHeader(http.StatusConflict)
		return
	} else {
		err := repo.bookTimeslot(ts)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
