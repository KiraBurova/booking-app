package timeslots

type Timeslot struct {
	OwnerId    string `json:"ownerId"`
	BookedById string `json:"bookedById"`
	Time       string `json:"time"`
	Booked     bool   `json:"booked"`
}
