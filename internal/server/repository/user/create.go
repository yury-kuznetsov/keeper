package user

import (
	"context"

	"github.com/google/uuid"
)

// Create inserts a new user entry into the database with the provided email and password.
// It returns the UUID of the newly created user or an error if the insertion fails.
func (r *Repository) Create(ctx context.Context, email, password string) (uuid.UUID, error) {
	id := uuid.New()
	query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`
	if _, err := r.db.ExecContext(ctx, query, id, email, password); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
