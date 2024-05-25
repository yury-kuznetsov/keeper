package data

import (
	"bytes"
	"context"
	"errors"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Push adds or updates a piece of data with the specified ID, UserID, data, and version number to the repository.
// It returns the updated version number and an error if any occurs.
func (s *Service) Push(ctx context.Context, id, userID uuid.UUID, category int, data []byte, version int) (int, error) {
	dataOld, _ := s.r.Find(ctx, id, userID)
	if dataOld.ID != uuid.Nil {
		if version != dataOld.Version {
			return version, errors.New("incorrect version")
		}
		if bytes.Equal(data, dataOld.Data) {
			return version, nil
		}
	}

	version++
	entity := model.Data{
		ID:       id,
		UserID:   userID,
		Category: category,
		Data:     data,
		Version:  version,
	}

	return version, s.r.Save(ctx, entity)
}
