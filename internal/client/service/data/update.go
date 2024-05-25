package data

import (
	"bytes"
	"errors"

	"github.com/google/uuid"
)

// Update updates the data of a specific entity identified by the given ID.
// It retrieves the entity from the repository using the ID.
// If the entity does not exist (ID is uuid.Nil), it returns an error with the message "entity not found".
// If the data provided is equal to the data in the entity, it returns nil, indicating no update is necessary.
// If the data provided is different from the data in the entity, it updates the entity's data field with the new data.
// It then saves the updated entity using the repository.
//
// Parameters:
// - id: the UUID of the entity to be updated
// - data: the new data to be assigned to the entity
//
// Returns:
// - error: an error object if there was any issue updating the entity or saving it
func (s *Service) Update(id uuid.UUID, data []byte) error {
	entity := s.r.Find(id)
	if entity.ID == uuid.Nil {
		return errors.New("entity not found")
	}

	if bytes.Equal(entity.Data, data) {
		return nil
	}

	entity.Data = data

	return s.r.Save(entity)
}
