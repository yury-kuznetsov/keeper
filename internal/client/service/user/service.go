package user

// Client is an interface that defines the behavior of a client.
type Client interface {
	Register(email, password string) error
	Login(email, password string) error
}

// Service is a struct that represents a service.
//
// It has the following field:
// - c: a Client that defines the behavior of a client.
type Service struct {
	c Client
}

// NewService creates a new Service instance.
//
// Parameters:
// - c: a Client that defines the behavior of a client.
//
// Returns:
// - *Service: a pointer to the newly created Service instance.
func NewService(c Client) *Service {
	return &Service{c: c}
}
