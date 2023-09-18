package timeslots

type Timeslot struct {
	// the person who books the timeslot
	UserId string `json:"userId"`
	// the person in whose calendar timeslot is booked
	InvitedUserId string `json:"InvitedUserId"`
	Time          string `json:"Time"`
	Booked        bool   `json:"booked"`
}
