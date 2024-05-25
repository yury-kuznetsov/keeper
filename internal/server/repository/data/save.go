package data

import (
	"context"
	"gophkeeper/internal/server/model"
)

// Save updates or inserts a piece of data into the repository. The data is represented
// by the provided Data object. The method executes an SQL query to either insert a new
// row into the 'data' table or update an existing row based on the ID.
// It returns an error if the execution of the query fails.
func (r *Repository) Save(ctx context.Context, data model.Data) error {
	query := `
		INSERT INTO data (id, user_id, category, data, version) VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE SET user_id = $2, category = $3, data = $4, version = $5
	`
	_, err := r.db.ExecContext(
		ctx,
		query,
		data.ID, data.UserID, data.Category, data.Data, data.Version,
	)

	return err
}
