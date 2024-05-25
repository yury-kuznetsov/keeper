package data

import (
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// Create generates a new UUID, creates a model.Data object with the provided category and data,
// and saves it using the underlying repository. It returns the newly generated UUID and any error occurred during the save operation.
//
// Parameters:
// - category: an integer representing the category of the data
// - data: a byte slice containing the data to be saved
//
// Returns:
// - uuid.UUID: the newly generated UUID
// - error: an error object if there was any issue saving the data
func (s *Service) Create(category int, data []byte) (uuid.UUID, error) {
	id := uuid.New()

	return id, s.r.Save(model.Data{
		ID:       id,
		Category: category,
		Data:     data,
		Version:  0,
	})
}
