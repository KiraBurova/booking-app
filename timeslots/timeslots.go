package timeslots

type Timeslot struct {
	CreatorId     string `json:"CreatorId"`
	InvitedUserId string `json:"InvitedUserId"`
	Time          string `json:"Time"`
	Booked        bool   `json:"booked"`
}
