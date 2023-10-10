package timeslots

import (
	"encoding/json"
	"net/http"
	"timezone-converter/db"

	"github.com/google/go-cmp/cmp"
)

type TimeslotData struct {
	OwnerId string       `json:"ownerId"`
	Time    []TimePeriod `json:"time"`
}

// time payload format that works
// "time": [{"From": "2014-11-12T11:45:26.371Z", "To": "2017-11-12T11:45:26.371Z"}]
func CreateTimeslots(w http.ResponseWriter, r *http.Request) {
	var timeslotsData TimeslotData
	json.NewDecoder(r.Body).Decode(&timeslotsData)

	repo := NewRepository(db.DbInstance)

	for _, timePeriod := range timeslotsData.Time {
		if !isTimePeriodValid(timePeriod) {
			return
		}
	}

	for i := 0; i < len(timeslotsData.Time); i++ {
		jsonTime, _ := json.Marshal(timeslotsData.Time[i])
		t := Timeslot{OwnerId: timeslotsData.OwnerId, Time: jsonTime, Booked: false}

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

	// TODO: data.Time is empty here
	ts, err := repo.getTimeslot(data)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var foundTimeslot Timeslot

	for i := 0; i < len(ts); i++ {
		var time TimePeriod
		marshalError := json.Unmarshal(ts[i].Time, &time)

		if marshalError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if cmp.Equal(time, data.Time) {
			foundTimeslot = ts[i]
		}
	}

	if foundTimeslot.Booked {
		w.WriteHeader(http.StatusConflict)
		return
	} else {
		err := repo.bookTimeslot(foundTimeslot)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
