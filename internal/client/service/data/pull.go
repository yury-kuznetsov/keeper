package data

import (
	"gophkeeper/internal/client/model"

	"github.com/google/uuid"
)

// Pull retrieves data from a remote source using the given ID and saves it locally.
// It returns an error if there was an issue pulling the data or saving it locally.
func (s *Service) Pull(id uuid.UUID) error {
	remote, err := s.c.Pull(id)
	if err != nil {
		return err
	}

	entity := model.Data{
		ID:       id,
		Category: remote.Category,
		Data:     remote.Data,
		Version:  remote.Version,
	}

	return s.r.Save(entity)
}
