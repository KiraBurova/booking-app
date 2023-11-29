package timeslots

import (
	"encoding/json"
	"net/http"
	"timezone-converter/db"
)

// time payload format that works
// "time": [{"From": "2014-11-12T11:45:26.371Z", "To": "2017-11-12T11:45:26.371Z"}]
func CreateTimeslots(w http.ResponseWriter, r *http.Request) {
	var timeslotsData TimeslotData
	json.NewDecoder(r.Body).Decode(&timeslotsData)

	// check if BookingDay was not send
	if timeslotsData.BookingDay.IsZero() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	repo := NewRepository(db.DbInstance)

	if !timeperiodsBelongToTheDay(timeslotsData.Time, timeslotsData.BookingDay) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if areTimePeriodsOverlapping(timeslotsData.Time) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, timePeriod := range timeslotsData.Time {
		if !isTimePeriodValid(timePeriod) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	err := repo.createTimeslots(timeslotsData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type BookTimeslotData struct {
	Id         string `json:"id"`
	BookedById string `json:"bookedById"`
}

func BookTimeslot(w http.ResponseWriter, r *http.Request) {
	repo := NewRepository(db.DbInstance)

	var data BookTimeslotData
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
