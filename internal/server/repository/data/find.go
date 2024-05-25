package data

import (
	"context"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Find retrieves a piece of data from the repository based on the provided ID and UserID.
// It returns a Data object containing the retrieved information and any error encountered during the retrieval process.
func (r *Repository) Find(ctx context.Context, id, userID uuid.UUID) (model.Data, error) {
	var data model.Data
	err := r.db.QueryRowContext(
		ctx,
		`SELECT id, user_id, category, data, version FROM data WHERE id = $1 AND user_id = $2`,
		id, userID,
	).Scan(&data.ID, &data.UserID, &data.Category, &data.Data, &data.Version)

	return data, err
}
