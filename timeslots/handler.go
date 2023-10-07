package timeslots

import (
	"encoding/json"
	"log"
	"net/http"
	"timezone-converter/db"
)

type TimeslotData struct {
	OwnerId string       `json:"ownerId"`
	Time    []TimePeriod `json:"time"`
}

func CreateTimeslots(w http.ResponseWriter, r *http.Request) {
	var timeslotsData TimeslotData
	json.NewDecoder(r.Body).Decode(&timeslotsData)

	repo := NewRepository(db.DbInstance)

	log.Println(timeslotsData.Time)

	timeslot := Timeslot{OwnerId: timeslotsData.OwnerId, Time: json.Marshal(timeslotsData.Time), Booked: false}

	err := repo.createTimeslots(timeslot)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
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
