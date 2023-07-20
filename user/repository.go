package user

import (
	"database/sql"
	"log"
	"timezone-converter/db"

	_ "github.com/mattn/go-sqlite3"
)

// type Repository interface {
// 	Create(user User) (int, error)
// 	Update(id int, user User) (User, error)
// 	Delete(id int) error
// 	UserExists(id int) (User, error)
// 	GetById(id int) (User, error)
// }

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	if db == nil {
		log.Panic("DB was not created.")
	}
	return &Repository{db: db}
}

func (r Repository) Create(user User) error {
	query := "INSERT INTO users(username, password, timeslots) values(?,?,?)"

	_, err := db.DbInstance.Exec(query, user.Username, user.Password, user.Timeslots)

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

func (r Repository) Update(id int, user User)     {}
func (r Repository) Delete(id int)                {}
func (r Repository) UserExists(id int, user User) {}
