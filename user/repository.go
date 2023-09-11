package user

import (
	"database/sql"
	"errors"
	"log"
	"timezone-converter/db"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func createTimeslotsTable() {
	const create = `CREATE TABLE IF NOT EXISTS timeslots(creatorId TEXT, invitedUserId TEXT, time TEXT, booked INTEGER)`

	if _, err := db.DbInstance.Exec(create); err != nil {
		log.Fatal(err)
	}
}

func (r Repository) CreateTimeslots(timeslot Timeslot) error {
	createTimeslotsTable()

	query := "INSERT INTO timeslots(creatorId, InvitedUserId, time, booked) values(?,?,?,true)"

	_, err := db.DbInstance.Exec(query, timeslot.CreatorId, timeslot.InvitedUserId, timeslot.Time, timeslot.Booked)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) isTimeslotBooked(timeslot Timeslot) (Timeslot, error) {
	ts := Timeslot{}
	query := "SELECT * FROM timeslots WHERE time=$1 AND invitedUserId=$2"

	row := db.DbInstance.QueryRow(query, timeslot.Time, timeslot.InvitedUserId)

	err := row.Scan(&ts.CreatorId, &ts.InvitedUserId, &ts.Time, &ts.Booked)

	if err != nil {
		return ts, err
	}

	return ts, nil
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

func (r Repository) UserExists(username string) (bool, error) {
	user, err := r.GetByUsername(username)

	if user.Username == username {
		return true, errors.New("User already exists")
	}

	return false, err
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
