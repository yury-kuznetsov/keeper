package user

// Login is a method of the Service struct that allows a user to login using their email and password.
// It calls the Login method of the underlying Client interface.
// Parameters:
// - email: the email of the user
// - password: the password of the user
// Returns:
// - error: returns an error if the login process fails.
func (s *Service) Login(email, password string) error {
	return s.c.Login(email, password)
}
