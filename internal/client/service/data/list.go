package data

import (
	"gophkeeper/internal/client/model"
)

// List returns a list of model.Data records by calling the FindAll method of the underlying repository.
// It returns the list of records and an error if there was an issue retrieving the data.
func (s *Service) List() ([]model.Data, error) {
	return s.r.FindAll()
}
