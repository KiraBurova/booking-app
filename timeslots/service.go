package timeslots

import (
	"timezone-converter/db"
)

func ServiceCreateTimeslots(timeslot TimeslotData) error {
	repo := NewRepository(db.DbInstance)
	bookingDayInUnix := timeslot.BookingDay.Unix()

	count, err := repo.getTimeslotsCountByBookingDay(bookingDayInUnix)

	if err != nil {
		return err
	}

	if count > 0 {
		query := "DELETE FROM timeslots WHERE BookingDay=?"

		_, err := db.DbInstance.Exec(query, bookingDayInUnix)

		if err != nil {
			return err
		}
	}

	err = insertTimeslotIntoDB(timeslot)

	if err != nil {
		return err
	}

	return nil
}
