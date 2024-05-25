package user

import (
	"database/sql"
)

// Repository is a type that represents a user storage object
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new instance of the Repository struct with the provided *sql.DB.
// It also executes a SQL statement to create the "users" table if it doesn't exist.
func NewRepository(db *sql.DB) *Repository {
	_, _ = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id       uuid    not null constraint users_pk primary key,
		email    varchar not null constraint users_pk_2 unique,
		password varchar not null
	)`)

	return &Repository{db: db}
}
