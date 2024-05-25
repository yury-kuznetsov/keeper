package user

// Service represents a service that provides user registration and login functionality.
// The methods Register and Login are used to handle user registration and login, respectively.
type Service interface {
	Register(email, password string) error
	Login(email, password string) error
}
