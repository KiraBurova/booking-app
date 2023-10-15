package timeslots

import (
	"time"
)

type TimeslotBase struct {
	OwnerId    string `json:"ownerId"`
	BookedById string `json:"bookedById"`
	Booked     bool   `json:"booked"`
}

type Timeslot struct {
	TimeslotBase
	Time []TimePeriod `json:"time"`
}

type TimeslotInDb struct {
	TimeslotBase
	TimeFrom int64 `json:"timeFrom"`
	TimeTo   int64 `json:"timeTo"`
}

type TimePeriod struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func isTimePeriodValid(timeperiod TimePeriod) bool {
	return timeperiod.From.Before(timeperiod.To)
}
