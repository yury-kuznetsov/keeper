package data

import (
	"context"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Versions retrieves all data versions associated with a given user ID.
// It executes an SQL query to select the ID and Version from the 'data' table
// where the user ID matches the provided value. The retrieved data versions are
// returned as an array of DataVersion objects. An error is returned if the query
// execution fails.
// The method requires a context.Context object to be passed as the first parameter
// and a uuid.UUID object representing the user ID as the second parameter.
// It returns an array of model.DataVersion objects and an error.
func (r *Repository) Versions(ctx context.Context, userID uuid.UUID) ([]model.DataVersion, error) {
	rows, err := r.db.QueryContext(
		ctx,
		"SELECT id, version FROM data WHERE user_id = $1",
		userID,
	)
	if err != nil {
		return nil, err
	}

	var versions []model.DataVersion
	for rows.Next() {
		var version model.DataVersion
		err := rows.Scan(
			&version.ID,
			&version.Version,
		)
		if err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return versions, nil
}
