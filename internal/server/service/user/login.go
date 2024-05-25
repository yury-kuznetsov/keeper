package user

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Login is a method of the Service struct that handles user authentication. It takes in the following parameters:
// - ctx: the context.Context object
// - email: the email of the user
// - password: the password of the user
// It returns the following values:
// - uuid.UUID: the ID of the authenticated user
// - error: an error if the authentication fails
// The Login method first checks if a user with the given email exists. If not, it returns an error.
// It then compares the provided password with the stored password using bcrypt. If they don't match, it returns an error.
// If the authentication is successful, it returns the user ID.
func (s *Service) Login(ctx context.Context, email, password string) (uuid.UUID, error) {
	// проверяем наличие пользователя с таким email
	user, err := s.r.FindByEmail(ctx, email)
	if err != nil {
		return uuid.Nil, errors.New("incorrect email or password")
	}

	// проверяем пароль
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return uuid.Nil, errors.New("incorrect email or password")
	}

	return user.ID, err
}
