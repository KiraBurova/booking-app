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

func (r Repository) createTimeslot(timeslot Timeslot) error {
	createTimeslotsTable()

	timeInUnixFrom := timeslot.TimeFrom.Unix()
	timeInUnixTo := timeslot.TimeTo.Unix()

	t := TimeslotInDB{TimeslotBase: TimeslotBase{Id: timeslot.Id, OwnerId: timeslot.OwnerId, Booked: false}, TimeFrom: timeInUnixFrom, TimeTo: timeInUnixTo}

	query := "INSERT INTO timeslots(ownerId, bookedById, timeFrom, timeTo, booked) values(?,?,?,?,?,?)"

	_, err := db.DbInstance.Exec(query, t.Id, t.OwnerId, t.BookedById, t.TimeFrom, t.TimeTo, t.Booked)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) bookTimeslot(timeslot TimeslotInDB) error {
	query := `UPDATE timeslots SET booked = $1, bookedById = $2 WHERE id=$3 AND ownerId=$4`
	_, err := db.DbInstance.Exec(query, 1, timeslot.BookedById, timeslot.Id, timeslot.OwnerId)

	if err != nil {
		return err
	}

	return nil
}
