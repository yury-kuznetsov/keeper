package data

import (
	"flag"

	"github.com/google/uuid"
)

// PushHandler handles pushing data to the service with the provided UUID.
// It takes a Service interface and a slice of arguments as input.
// The function parses the provided arguments, extracts the UUID, and calls the Push method of the Service interface using the extracted UUID.
// It returns any error encountered during the push process.
//
// Example usage:
//
//	push -i <UUID>
func PushHandler(svc Service, arguments []string) error {
	fs := flag.NewFlagSet("pull", flag.ExitOnError)
	id := fs.String("i", "", "id")

	err := fs.Parse(arguments)
	if err != nil {
		return err
	}

	uid, err := uuid.Parse(*id)
	if err != nil {
		return err
	}

	return svc.Push(uid)
}
