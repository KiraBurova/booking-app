package user

import (
	"database/sql"
	"errors"
	"timezone-converter/db"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r Repository) Create(user User) error {
	query := "INSERT INTO users(id, username, password, timeslots) values(?,?,?,?)"

	_, err := db.DbInstance.Exec(query, user.Id, user.Username, user.Password, user.Timeslots)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetById(id string) (User, error) {
	user := User{}
	query := "SELECT * FROM users WHERE id=?"

	row := db.DbInstance.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Timeslots)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r Repository) GetByUsername(username string) (User, error) {
	user := User{}
	query := "SELECT * FROM users WHERE username=?"

	row := db.DbInstance.QueryRow(query, username)

	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Timeslots)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r Repository) UserExists(username string) (bool, error) {
	user, err := r.GetByUsername(username)

	if user.Username == username {
		return true, errors.New("User already exists")
	}

	return false, err
}

func (r Repository) Update(id int, user User) {}
func (r Repository) Delete(id int)            {}
