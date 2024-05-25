package user

import (
	"flag"
)

// LoginHandler parses the provided email and password flags and calls the Login method of the provided Service.
// It returns an error if the login fails.
//
// Example usage:
//
//	login -u user@email.com -p 12345678
func LoginHandler(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("register", flag.ExitOnError)
	email := fs.String("u", "", "Your email")
	password := fs.String("p", "", "Your password")
	_ = fs.Parse(arguments)

	return svc.Login(*email, *password)
}
