package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY AUTOINCREMENT, username TEXT, password TEXT, timeslots BLOB);`

var DbInstance *sql.DB

func InitDb() {

	db, err := sql.Open("sqlite3", "./users.db")

	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(create); err != nil {
		log.Fatal(err)
	}

	DbInstance = db
}
