package timeslots

import (
	"encoding/json"
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

	for i := 0; i < len(timeslotsData.Time); i++ {
		t := Timeslot{OwnerId: timeslotsData.OwnerId, Time: timeslotsData.Time[i], Booked: false}

		err := repo.createTimeslots(t)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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
