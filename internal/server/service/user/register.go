package user

import (
	"context"
	"errors"
	"net/mail"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Register is a method of the Service struct that handles user registration. It takes in the following parameters:
// - ctx: the context.Context object
// - email: the email of the user to be registered
// - password: the password of the user to be registered
// It returns the following values:
// - uuid.UUID: the ID of the registered user
// - error: an error if the registration fails
// The Register method first validates the email and password using the `validate` function.
// If the validation fails, it returns an error.
// It then checks if a user with the given email already exists. If so, it returns an error.
// The password is hashed using bcrypt and then stored.
// Finally, it registers the user in the repository and returns the user ID.
func (s *Service) Register(ctx context.Context, email, password string) (uuid.UUID, error) {
	if err := validate(email, password); err != nil {
		return uuid.Nil, err
	}

	// проверяем наличие пользователя с таким email
	user, _ := s.r.FindByEmail(ctx, email)
	if user.ID != uuid.Nil {
		return uuid.Nil, errors.New("user already exists")
	}

	// подготавливаем пароль к хранению
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	// регистрируем пользователя
	userID, err := s.r.Create(ctx, email, string(passwordHash))
	if err != nil {
		return uuid.Nil, err
	}

	return userID, err
}

func validate(email, password string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("incorrect email")
	}
	if len(password) < 8 {
		return errors.New("password is too short")
	}
	if len(password) > 72 {
		return errors.New("password is too long")
	}
	return nil
}
