package user

import (
	"context"

	"github.com/google/uuid"
)

// Service represents a service that provides user registration and login functionality.
type Service interface {
	Register(ctx context.Context, email, password string) (uuid.UUID, error)
	Login(ctx context.Context, email, password string) (uuid.UUID, error)
}

// JWTService represents a service that provides functionality for generating JWT tokens.
type JWTService interface {
	GenerateToken(userID uuid.UUID) string
}
