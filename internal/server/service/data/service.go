package data

import (
	"context"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Repository is an interface that defines methods for interacting with a data storage system.
type Repository interface {
	Versions(ctx context.Context, userID uuid.UUID) ([]model.DataVersion, error)
	Find(ctx context.Context, id, userID uuid.UUID) (model.Data, error)
	Save(ctx context.Context, data model.Data) error
}

// Service is a type that represents a data service object
type Service struct {
	r Repository
}

// NewService returns a new instance of the Service struct with the provided Repository interface.
// It initializes the Service struct with the provided Repository and returns a pointer to it.
func NewService(r Repository) *Service {
	return &Service{r: r}
}
