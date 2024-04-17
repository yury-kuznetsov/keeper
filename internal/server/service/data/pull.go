package data

import (
	"context"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Pull retrieves a piece of data with the specified ID and UserID from the repository.
// It returns the data and an error if any occurs.
func (s *Service) Pull(ctx context.Context, id, userID uuid.UUID) (model.Data, error) {
	return s.r.Find(ctx, id, userID)
}
