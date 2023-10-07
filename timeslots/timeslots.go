package timeslots

import "time"

type Timeslot struct {
	OwnerId    string       `json:"ownerId"`
	BookedById string       `json:"bookedById"`
	Time       []TimePeriod `json:"time"`
	Booked     bool         `json:"booked"`
}

type TimePeriod struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func isTimePeriodValid(timeperiod TimePeriod) {
	return timeperiod.From.Before(timeperiod.To)
}