package user

import (
	"flag"
)

// RegisterHandler handles the registration of a new user by parsing the provided email and password flags,
// and calling the Register method of the provided Service. It returns an error if the registration fails.
//
// Example usage:
//
//	register -u user@email.com -p 12345678
func RegisterHandler(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("register", flag.ExitOnError)
	email := fs.String("u", "", "Your email")
	password := fs.String("p", "", "Your password")
	_ = fs.Parse(arguments)

	return svc.Register(*email, *password)
}
