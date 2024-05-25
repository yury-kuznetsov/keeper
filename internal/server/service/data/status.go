package data

import (
	"context"
	"gophkeeper/internal/server/model"

	"github.com/google/uuid"
)

// Status returns the list of versions of data associated with the specified userID.
//
// It takes a context object and a userID as arguments and returns a slice of DataVersion structs
// and an error. Each DataVersion struct consists of an ID (uuid.UUID) and a Version number (int).
// The method calls the Versions method on the underlying Repository to retrieve the data versions.
//
// Example usage:
//
//	svc := &Service{r: repository}
//	userID := uuid.New()
//	versions, err := svc.Status(context.Background(), userID)
func (s *Service) Status(ctx context.Context, userID uuid.UUID) ([]model.DataVersion, error) {
	return s.r.Versions(ctx, userID)
}
