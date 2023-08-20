package auth

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

func createTable() {
	const create = `CREATE TABLE IF NOT EXISTS sessions(token TEXT, userId TEXT)`

	if _, err := db.DbInstance.Exec(create); err != nil {
		log.Fatal(err)
	}
}

func (r Repository) Create(token string, userId string) error {
	createTable()

	log.Println(token, userId)

	query := "INSERT INTO sessions(token, userId) values(?,?)"

	_, err := db.DbInstance.Exec(query, token, userId)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Delete(token string) error {
	query := "DELETE FROM sessions WHERE token=?"

	_, err := db.DbInstance.Exec(query, token)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Get(token string) (bool, error) {
	var t string
	query := "SELECT * FROM sessions WHERE token=?"

	row := db.DbInstance.QueryRow(query)

	err := row.Scan(&t)

	return t == token, err
}
