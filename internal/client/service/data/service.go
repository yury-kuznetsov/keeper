package data

import (
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// Repository is an interface for performing CRUD operations on a data store.
type Repository interface {
	Find(id uuid.UUID) model.Data
	FindAll() ([]model.Data, error)
	Save(data model.Data) error
}

// Client is an interface for interacting with a remote data store.
type Client interface {
	Pull(id uuid.UUID) (model.Data, error)
	Push(model.Data) (int, error)
	Status() ([]model.DataVersionRemote, error)
}

// Service is a type that represents a service which performs operations on data using a repository and a client.
type Service struct {
	r Repository
	c Client
}

// NewService creates a new instance of the Service type with the provided Repository and Client
func NewService(r Repository, c Client) *Service {
	return &Service{r: r, c: c}
}
