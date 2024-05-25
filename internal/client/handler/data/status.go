package data

import (
	"fmt"
)

// StatusHandler handles the retrieval and printing of status information
// for a given service. It takes a Service interface as input and calls its
// Status method to retrieve the status models. It then iterates over the
// models and prints the ID, VersionLocal, and VersionRemote for each model.
// If an error occurs during the retrieval of the models, it is returned.
// Otherwise, it returns nil.
func StatusHandler(svc Service) error {
	var err error

	models, err := svc.Status()
	if err != nil {
		return err
	}

	for _, model := range models {
		fmt.Println("---")
		fmt.Printf("ID: %s \n", model.ID)
		fmt.Printf("VersionLocal: %d \n", model.VersionLocal)
		fmt.Printf("VersionRemote: %d \n", model.VersionRemote)
	}

	return nil
}
