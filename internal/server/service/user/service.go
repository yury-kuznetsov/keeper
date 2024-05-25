package user

import (
	"context"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Repository is an interface that defines methods for interacting with a user storage system.
type Repository interface {
	Create(ctx context.Context, email, password string) (uuid.UUID, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
}

// Service is a type that represents a user service object
type Service struct {
	r Repository
}

// NewService creates a new instance of the Service struct with the provided Repository interface.
// It initializes the Service struct with the provided Repository and returns a pointer to it.
func NewService(r Repository) *Service {
	return &Service{r: r}
}
