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
	const create = `CREATE TABLE IF NOT EXISTS timeslots(ownerId TEXT, bookedById TEXT, timeFrom INTEGER, timeTo INTEGER, booked INTEGER)`

	if _, err := db.DbInstance.Exec(create); err != nil {
		log.Fatal(err)
	}
}

func (r Repository) createTimeslot(timeslot TimeslotInDB) error {
	createTimeslotsTable()

	query := "INSERT INTO timeslots(ownerId, bookedById, timeFrom, timeTo, booked) values(?,?,?,?,?)"

	_, err := db.DbInstance.Exec(query, timeslot.OwnerId, timeslot.BookedById, timeslot.TimeFrom, timeslot.TimeTo, timeslot.Booked)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) getTimeslot(timeslot Timeslot) (TimeslotInDB, error) {

	ts := TimeslotInDB{}

	for i := 0; i < len(timeslot.Time); i++ {
		timeInUnixTo := timeslot.Time[i].To.Unix()
		timeInUnixFrom := timeslot.Time[i].From.Unix()

		query := "SELECT * FROM timeslots WHERE ownerId=$2 AND timeFrom=$3 AND timeTo=$4"

		row := db.DbInstance.QueryRow(query, timeslot.OwnerId, timeInUnixFrom, timeInUnixTo)

		err := row.Scan(&ts.OwnerId, &ts.BookedById, &ts.TimeFrom, &ts.TimeTo, &ts.Booked)

		if err != nil {
			return ts, err
		}
	}

	return ts, nil

}

func (r Repository) bookTimeslot(timeslot TimeslotInDB) error {
	query := `UPDATE timeslots SET booked = $1, bookedById = $2 WHERE timeFrom=$3 AND timeTo=$4 AND ownerId=$5`
	_, err := db.DbInstance.Exec(query, 1, timeslot.BookedById, timeslot.TimeFrom, timeslot.TimeTo, timeslot.OwnerId)

	if err != nil {
		return err
	}

	return nil
}
