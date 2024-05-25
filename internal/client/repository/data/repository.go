package data

import (
	"database/sql"
)

// Repository структура репозитория данных
type Repository struct {
	db *sql.DB
}

// NewRepository creates a new Repository instance with the provided database connection.
//
// It executes a CREATE TABLE IF NOT EXISTS query to create a "data" table in the
// database if it doesn't already exist.
// The error encountered during the execution is ignored.
// The created Repository instance is returned.
func NewRepository(db *sql.DB) *Repository {
	_, _ = db.Exec(`
	CREATE TABLE IF NOT EXISTS data (
		id       uuid     not null primary key,
		category smallint not null,
		data     text,
		version  integer  not null
	)`)

	return &Repository{db: db}
}
