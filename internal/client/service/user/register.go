package user

import (
	"errors"
	"net/mail"
)

// Register is a method of the Service struct that allows a user to register using their email and password.
// It calls the Register method of the underlying Client interface after validating the email and password.
// Parameters:
// - email: the email of the user
// - password: the password of the user
// Returns:
// - error: returns an error if the registration process fails.
func (s *Service) Register(email, password string) error {
	if err := validate(email, password); err != nil {
		return err
	}
	return s.c.Register(email, password)
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
