package data

import (
	"gophkeeper/internal/client/model"
)

// Save inserts or updates a data record in the repository.
//
// It executes an INSERT INTO query to insert a new record into the "data" table.
// If a record with the same ID already exists, the query updates the record
// with the new values for category, data, and version.
// The data.ID, data.Category, data.Data, and data.Version values are used as
// parameters for the query.
// The error encountered during the execution is returned.
func (r *Repository) Save(data model.Data) error {
	query := `
		INSERT INTO data (id, category, data, version) VALUES ($1, $2, $3, $4)
		ON CONFLICT (id) DO UPDATE SET category = $2, data = $3, version = $4
	`
	_, err := r.db.Exec(
		query,
		data.ID, data.Category, data.Data, data.Version,
	)

	return err
}
