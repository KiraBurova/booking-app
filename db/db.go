package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const create string = `CREATE TABLE IF NOT EXISTS users(id TEXT, username TEXT, password TEXT, timeslots BLOB);`

var DbInstance *sql.DB

func InitDb(dbName string) {

	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s", dbName))

	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(create); err != nil {
		log.Fatal(err)
	}

	DbInstance = db
}
