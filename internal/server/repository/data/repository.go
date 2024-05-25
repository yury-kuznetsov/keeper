package data

import (
	"database/sql"
)

// Repository is a type that represents a data storage object
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new instance of the Repository struct with the provided *sql.DB.
// It also executes a SQL statement to create the "data" table if it doesn't exist.
func NewRepository(db *sql.DB) *Repository {
	_, _ = db.Exec(`
	CREATE TABLE IF NOT EXISTS data (
		id       uuid     not null constraint data_pk primary key,
		user_id  uuid     not null constraint data_users_id_fk references users,
		category smallint not null,
		data     text,
		version  int      not null
	)`)

	return &Repository{db: db}
}
