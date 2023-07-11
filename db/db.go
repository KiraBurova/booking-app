package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `create table users(id integer primary key autoincrement, username text, password text, timeslots blob);`

var DbInstance *sql.DB

func InitDb() error {
	db, err := sql.Open("sqlite3", "./db/users.db")
	if err != nil {
		return err
	}
	if _, err := db.Exec(create); err != nil {
		return err
	}
	DbInstance = db
	return nil
}
