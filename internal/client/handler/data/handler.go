package data

import (
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// Service is an interface that defines operations to interact with a service.
type Service interface {
	Create(category int, data []byte) (uuid.UUID, error)
	Update(id uuid.UUID, data []byte) error
	Remove(id uuid.UUID) error
	List() ([]model.Data, error)
	Pull(id uuid.UUID) error
	Push(id uuid.UUID) error
	Status() ([]model.DataVersion, error)
}
