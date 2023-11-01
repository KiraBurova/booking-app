package timeslots

import (
	"sort"
	"time"
)

type TimeslotBase struct {
	Id         string `json:"id"`
	OwnerId    string `json:"ownerId"`
	BookedById string `json:"bookedById"`
	Booked     bool   `json:"booked"`
}

type TimeslotInDB struct {
	TimeslotBase
	TimeFrom   int64 `json:"timeFrom"`
	TimeTo     int64 `json:"timeTo"`
	BookingDay int64 `json:"BookingDay"`
}

type TimeslotData struct {
	OwnerId    string       `json:"ownerId"`
	Time       []TimePeriod `json:"time"`
	BookingDay time.Time    `json:"BookingDay"`
}

type TimePeriod struct {
	From time.Time `json:"from"`
	To   time.Time `json:"to"`
}

func isTimePeriodValid(timeperiod TimePeriod) bool {
	return timeperiod.From.Before(timeperiod.To)
}

func areTimePeriodsNotOverlapping(timeperiods []TimePeriod) bool {
	sort.Slice(timeperiods, func(i, j int) bool {
		return timeperiods[i].From.Before(timeperiods[j].From)
	})

	prev := timeperiods[0]
	for i := 1; i < len(timeperiods); i++ {
		cur := timeperiods[i]
		if !prev.To.Before(cur.From) {
			return false
		}
		prev = cur
	}
	return true
}

func timeperiodsBelongToTheDay(timeperiods []TimePeriod, day time.Time) bool {
	for i := 0; i < len(timeperiods); i++ {
		period := timeperiods[i]
		fromBelongsToDay := period.From.Day() == day.Day() && period.From.Month() == day.Month() && period.From.Year() == day.Year()
		toBelongsToDay := period.To.Day() == day.Day() && period.To.Month() == day.Month() && period.To.Year() == day.Year()

		if !(fromBelongsToDay && toBelongsToDay) {
			return false
		}
	}
	return true
}
