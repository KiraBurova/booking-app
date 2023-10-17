package timeslots

import (
	"sort"
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

func sortByTimeFrom(timeperiods []TimePeriod) {
	sort.Slice(timeperiods, func(i, j int) bool {
		return timeperiods[i].From.Before(timeperiods[j].From)
	})
}

func areTimePeriodsOverlapping(timeperiods []TimePeriod) bool {

	// sort timeperiods from earliest "from"

	// check that the next "from" is bigger than the previous
	// check that previous "to" is smaller than next "from"
	// from 16th to 18th
	// from 19th to 20th

	// from 17th to 18th
	// from 17th to 19th

	return false
}
