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
	const create = `CREATE TABLE IF NOT EXISTS timeslots(ownerId TEXT, bookedById TEXT, time BLOB, booked INTEGER)`

	if _, err := db.DbInstance.Exec(create); err != nil {
		log.Fatal(err)
	}
}

func (r Repository) createTimeslots(timeslot Timeslot) error {
	createTimeslotsTable()

	query := "INSERT INTO timeslots(ownerId, bookedById, time, booked) values(?,?,?,?)"

	// TODO: error  sql: converting argument $3 type: unsupported type timeslots.TimePeriod, a struct
	_, err := db.DbInstance.Exec(query, timeslot.OwnerId, timeslot.BookedById, timeslot.Time, timeslot.Booked)

	log.Println(err)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) getTimeslot(timeslot Timeslot) (Timeslot, error) {
	ts := Timeslot{}
	query := "SELECT * FROM timeslots WHERE time=$1 AND ownerId=$2"

	row := db.DbInstance.QueryRow(query, timeslot.Time, timeslot.BookedById)

	err := row.Scan(&ts.OwnerId, &ts.BookedById, &ts.Time, &ts.Booked)

	if err != nil {
		return ts, err
	}

	return ts, nil
}

func (r Repository) bookTimeslot(timeslot Timeslot) error {
	query := `UPDATE timeslots SET booked = $1, bookedById = $2 WHERE time=$3 AND ownerId=$4`
	_, err := db.DbInstance.Exec(query, 1, timeslot.BookedById, timeslot.Time, timeslot.OwnerId)

	if err != nil {
		return err
	}

	return nil
}
