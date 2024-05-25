package data

import (
	"context"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Service is an interface that defines a set of methods for managing data.
type Service interface {
	Pull(ctx context.Context, id, userID uuid.UUID) (model.Data, error)
	Push(ctx context.Context, id, userID uuid.UUID, category int, data []byte, version int) (int, error)
	Status(ctx context.Context, userID uuid.UUID) ([]model.DataVersion, error)
}
