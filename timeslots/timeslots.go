package timeslots

type Timeslot struct {
	OwnerId    string     `json:"ownerId"`
	BookedById string     `json:"bookedById"`
	Time       TimePeriod `json:"time"`
	Booked     bool       `json:"booked"`
}

type TimePeriod struct {
	From string `json:"from"`
	To   string `json:"to"`
}
