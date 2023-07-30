package auth

import (
	"database/sql"
	"timezone-converter/db"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Create(token string) error {
	query := "INSERT INTO sessions(token) values(?)"

	_, err := db.DbInstance.Exec(query, token)

	if err != nil {
		return err
	}

	return nil
}
