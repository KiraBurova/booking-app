package user

import (
	"database/sql"
	"errors"
	"io/fs"
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
	query := "INSERT INTO users(id, username, password) values(?,?,?)"

	_, err := db.DbInstance.Exec(query, user.Id, user.Username, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) GetById(id string) (User, error) {
	user := User{}
	query := "SELECT * FROM users WHERE id=?"

	row := db.DbInstance.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r Repository) GetByUsername(username string) (User, error) {
	user := User{}
	query := "SELECT * FROM users WHERE username=?"

	row := db.DbInstance.QueryRow(query, username)

	err := row.Scan(&user.Id, &user.Username, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r Repository) UserExists(username string) bool {
	_, err := r.GetByUsername(username)

	return errors.Is(err, fs.ErrExist)
}

func (r Repository) Update(user User) error {
	query := `UPDATE users SET username = $1, password = $2, WHERE id = $3`
	_, err := db.DbInstance.Exec(query, user.Username, user.Password, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Delete(id int) {}
