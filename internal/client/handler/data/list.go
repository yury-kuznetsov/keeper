package data

import "fmt"

// ListHandler sends a request to the provided service to retrieve a list of models.
// It iterates through each model in the list and prints its ID, category, and data,
// excluding any models with nil data.
// If any error occurs during the retrieval process, it returns that error.
func ListHandler(svc Service) error {
	var err error

	models, err := svc.List()
	if err != nil {
		return err
	}

	for _, model := range models {
		// пропускаем удаленные записи
		if model.Data != nil {
			fmt.Println("---")
			fmt.Printf("ID: %s \n", model.ID)
			fmt.Printf("Category: %d \n", model.Category)
			fmt.Printf("Data: %s \n", model.Data)
		}
	}

	return nil
}
