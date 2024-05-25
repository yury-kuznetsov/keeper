package data

import (
	"flag"

	"github.com/google/uuid"
)

// PullHandler handles pulling data from the service by parsing the provided "id" flag
// and calling the Pull method of the provided Service. It returns an error if the process fails.
//
// Usage example: `pull -i [id]`
func PullHandler(svc Service, arguments []string) error {
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

	return svc.Pull(uid)
}
