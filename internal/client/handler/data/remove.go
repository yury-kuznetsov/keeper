package data

import (
	"flag"

	"github.com/google/uuid"
)

// RemoveHandler removes a record from the service based on the provided ID.
// It takes a Service interface and a slice of arguments as input.
// The function creates a new flag set and adds a flag for the ID.
// The function then parses the arguments and extracts the ID flag value.
// The ID value is then parsed into a UUID.
// The function calls the Remove method of the Service interface, passing the parsed UUID.
// If any error occurs during flag parsing or UUID parsing, the error is returned.
// Otherwise, the function returns the result of the Remove method.
func RemoveHandler(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("remove", flag.ExitOnError)
	id := fs.String("i", "", "id")

	err := fs.Parse(arguments)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(*id)
	if err != nil {
		return err
	}

	return svc.Remove(uid)
}
