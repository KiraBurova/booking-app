package timeslots

import (
	"database/sql"
	"log"
	"timezone-converter/db"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func createTimeslotsTable() {
	const create = `CREATE TABLE IF NOT EXISTS timeslots(creatorId TEXT, invitedUserId TEXT, time TEXT, booked INTEGER)`

	if _, err := db.DbInstance.Exec(create); err != nil {
		log.Fatal(err)
	}
}

func (r Repository) CreateTimeslots(timeslot Timeslot) error {
	createTimeslotsTable()

	query := "INSERT INTO timeslots(UserId, InvitedUserId, time, booked) values(?,?,?,?)"

	_, err := db.DbInstance.Exec(query, timeslot.UserId, timeslot.InvitedUserId, timeslot.Time, timeslot.Booked)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) isTimeslotBooked(timeslot Timeslot) (Timeslot, error) {
	ts := Timeslot{}
	query := "SELECT * FROM timeslots WHERE time=$1 AND invitedUserId=$2"

	row := db.DbInstance.QueryRow(query, timeslot.Time, timeslot.InvitedUserId)

	err := row.Scan(&ts.UserId, &ts.InvitedUserId, &ts.Time, &ts.Booked)

	if err != nil {
		return ts, err
	}

	return ts, nil
}
