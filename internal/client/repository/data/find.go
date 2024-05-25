package data

import (
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// Find retrieves a single data record from the repository based on the given ID.
// It executes a SELECT query using the provided ID to fetch the corresponding
// data record from the "data" table. The retrieved values for ID, Category, Data,
// and Version are assigned to a new model.Data struct and returned.
func (r *Repository) Find(id uuid.UUID) model.Data {
	var data model.Data
	_ = r.db.QueryRow(
		`SELECT id, category, data, version FROM data WHERE id = $1`,
		id,
	).Scan(&data.ID, &data.Category, &data.Data, &data.Version)

	return data
}
