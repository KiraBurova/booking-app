package user

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type ConnectDB struct {
	db *sql.DB
}

func (d *ConnectDB) exec(query string, args any) {
	d.db.Exec(query, args)
}

func (d *ConnectDB) queryRow(query string, args any) User {
	user := User{}

	row := d.db.QueryRow(query, args)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Timeslots)

	// TODO: Scan error on column index 3, name "timeslots": unsupported Scan, storing driver.Value type []uint8 into type *map[string]user.TimeslotStatus

	if err != nil {
		log.Panic(err)
	}

	return user
}
