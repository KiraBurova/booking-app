package timeslots

import (
	"encoding/json"
	"net/http"
	"time"
	"timezone-converter/db"

	"github.com/google/uuid"
)

type TimeslotData struct {
	OwnerId    string       `json:"ownerId"`
	Time       []TimePeriod `json:"time"`
	BookingDay time.Time    `json:"BookingDay"`
}

// time payload format that works
// "time": [{"From": "2014-11-12T11:45:26.371Z", "To": "2017-11-12T11:45:26.371Z"}]
func CreateTimeslots(w http.ResponseWriter, r *http.Request) {
	var timeslotsData TimeslotData
	json.NewDecoder(r.Body).Decode(&timeslotsData)

	// check if BookingDay was not send
	if timeslotsData.BookingDay.IsZero() {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	repo := NewRepository(db.DbInstance)

	if !timeperiodsBelongToTheDay(timeslotsData.Time, timeslotsData.BookingDay) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if areTimePeriodsOverlapping(timeslotsData.Time) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, timePeriod := range timeslotsData.Time {
		if !isTimePeriodValid(timePeriod) {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	// timeslots := []TimeslotData{timeslotsData}
	// err := repo.createTimeslots(timeslots...)

	for i := 0; i < len(timeslotsData.Time); i++ {

		t := Timeslot{TimeslotBase: TimeslotBase{Id: uuid.NewString(), OwnerId: timeslotsData.OwnerId, Booked: false}, TimeFrom: timeslotsData.Time[i].From, TimeTo: timeslotsData.Time[i].To}

		err := repo.createTimeslot(t)

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

	ts, err := repo.getTimeslotById(data.Id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ts.BookedById = data.BookedById

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
