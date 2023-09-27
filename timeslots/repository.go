package timeslots

import (
	"database/sql"
	"errors"
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
	const create = `CREATE TABLE IF NOT EXISTS timeslots(userId TEXT, BookedTimeslotFromId TEXT, time TEXT, booked INTEGER)`

	if _, err := db.DbInstance.Exec(create); err != nil {
		log.Fatal(err)
	}
}

func (r Repository) createTimeslots(timeslot Timeslot) error {
	createTimeslotsTable()

	query := "INSERT INTO timeslots(userId, bookedTimeslotFromId, time, booked) values(?,?,?,?)"

	_, err := db.DbInstance.Exec(query, timeslot.UserId, timeslot.BookedTimeslotFromId, timeslot.Time, timeslot.Booked)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) userHasTimeslots(userId string) bool {
	ts := Timeslot{}
	query := "SELECT * FROM timeslots WHERE userId=?"

	row := db.DbInstance.QueryRow(query, userId)

	err := row.Scan(&ts.UserId, &ts.BookedTimeslotFromId, &ts.Time, &ts.Booked)

	return !(err != nil && errors.Is(err, sql.ErrNoRows))
}

func (r Repository) getTimeslot(timeslot Timeslot) (Timeslot, error) {
	ts := Timeslot{}
	query := "SELECT * FROM timeslots WHERE time=$1 AND userId=$2"

	row := db.DbInstance.QueryRow(query, timeslot.Time, timeslot.BookedTimeslotFromId)

	err := row.Scan(&ts.UserId, &ts.BookedTimeslotFromId, &ts.Time, &ts.Booked)

	if err != nil {
		return ts, err
	}

	return ts, nil
}

func (r Repository) bookTimeslot(timeslot Timeslot) error {
	query := `UPDATE timeslots SET booked = $1, BookedTimeslotFromId = $2 WHERE time=$3 AND userId=$4`
	_, err := db.DbInstance.Exec(query, 1, timeslot.BookedTimeslotFromId, timeslot.Time, timeslot.UserId)

	if err != nil {
		return err
	}

	return nil
}
