package timeslots

import (
	"database/sql"
	"errors"
	"log"
	"timezone-converter/db"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func createTimeslotsTable() {
	const create = `CREATE TABLE IF NOT EXISTS timeslots(id TEXT, ownerId TEXT, bookedById TEXT, timeFrom INTEGER, timeTo INTEGER, bookingDay INTEGER, booked INTEGER)`

	if _, err := db.DbInstance.Exec(create); err != nil {
		log.Fatal(err)
	}
}

func insertTimeslotIntoDB(timeslot TimeslotData) error {
	id := uuid.NewString()
	bookingDayInUnix := timeslot.BookingDay.Unix()

	for i := 0; i < len(timeslot.Time); i++ {
		timeInUnixFrom := timeslot.Time[i].From.Unix()
		timeInUnixTo := timeslot.Time[i].To.Unix()

		t := TimeslotInDB{TimeslotBase: TimeslotBase{Id: id, OwnerId: timeslot.OwnerId, Booked: false}, TimeFrom: timeInUnixFrom, TimeTo: timeInUnixTo, BookingDay: bookingDayInUnix}

		query := "INSERT INTO timeslots(id, ownerId, bookedById, timeFrom, timeTo, bookingDay, booked) values(?,?,?,?,?,?,?)"

		_, err := db.DbInstance.Exec(query, t.Id, t.OwnerId, t.BookedById, t.TimeFrom, t.TimeTo, t.BookingDay, t.Booked)

		if err != nil {
			return err
		}
	}
	return nil
}

func (r Repository) createTimeslots(timeslot TimeslotData) error {
	createTimeslotsTable()

	bookingDayInUnix := timeslot.BookingDay.Unix()

	ts, err := r.getTimeslotByBookingDay(bookingDayInUnix)

	if !errors.Is(err, sql.ErrNoRows) {
		query := "DELETE FROM timeslots WHERE BookingDay=?"

		_, err := db.DbInstance.Exec(query, ts.BookingDay)

		if err != nil {
			return err
		}

		// no new variables on left side of :=compiler
		// err := insertTimeslotIntoDB(timeslot)

		insertTimeslotIntoDB(timeslot)

	} else {
		err := insertTimeslotIntoDB(timeslot)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r Repository) getTimeslotByBookingDay(day int64) (TimeslotInDB, error) {
	timeslot := TimeslotInDB{}
	query := "SELECT * FROM timeslots WHERE bookingDay=?"

	row := db.DbInstance.QueryRow(query, day)

	err := row.Scan(&timeslot.Id, &timeslot.OwnerId, &timeslot.BookedById, &timeslot.TimeFrom, &timeslot.TimeTo, &timeslot.BookingDay, &timeslot.Booked)

	if err != nil {
		return timeslot, err
	}

	return timeslot, nil
}

func (r Repository) getTimeslotById(id string) (TimeslotInDB, error) {
	timeslot := TimeslotInDB{}
	query := "SELECT * FROM timeslots WHERE id=?"

	row := db.DbInstance.QueryRow(query, id)

	err := row.Scan(&timeslot.Id, &timeslot.OwnerId, &timeslot.BookedById, &timeslot.TimeFrom, &timeslot.TimeTo, &timeslot.BookingDay, &timeslot.Booked)

	if err != nil {
		return timeslot, err
	}

	return timeslot, nil
}

func (r Repository) bookTimeslot(timeslot TimeslotInDB) error {
	query := `UPDATE timeslots SET booked = $1, bookedById = $2 WHERE id=$3 AND ownerId=$4`
	_, err := db.DbInstance.Exec(query, 1, timeslot.BookedById, timeslot.Id, timeslot.OwnerId)

	if err != nil {
		return err
	}

	return nil
}
