package timeslots

type Timeslot struct {
	UserId               string `json:"userId"`
	BookedTimeslotFromId string `json:"bookedTimeslotFromId"`
	Time                 string `json:"time"`
	Booked               bool   `json:"booked"`
}
