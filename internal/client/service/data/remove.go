package data

import (
	"errors"

	"github.com/google/uuid"
)

// Remove removes the data associated with the given ID from the Service's repository.
// It sets the data field of the entity to nil and increments the version by 1.
// It returns an error if the entity is not found or if there was an issue saving the changes.
func (s *Service) Remove(id uuid.UUID) error {
	entity := s.r.Find(id)
	if entity.ID == uuid.Nil {
		return errors.New("entity not found")
	}

	entity.Data = nil
	entity.Version++

	return s.r.Save(entity)
}
