package api

import (
	"encoding/json"
	"gophkeeper/internal/client/model"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type statusResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
}

// Status sends a GET request to the server to retrieve the status of data versions.
// It prepares the request, sends it, and handles possible errors.
// If the request is successful, it returns an array of model.DataVersionRemote containing the ID and version of each data record.
// The Status method returns an error if the request fails.
//
// Example:
//
//	client := &Client{address: "https://example.com"}
//	status, err := client.Status()
//	if err != nil {
//	    log.Fatal(err)
//	}
//	for _, s := range status {
//	    fmt.Println("ID:", s.ID, "Version:", s.Version)
//	}
func (c *Client) Status() ([]model.DataVersionRemote, error) {
	path := "/api/data/status"

	// отправляем запрос
	req, err := http.NewRequest("GET", c.address+path, nil)
	if err != nil {
		return nil, err
	}

	// добавляем заголовок с JWT-токеном
	req.Header.Add("Authorization", "Bearer "+loadToken())

	// отсылаем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// извлекаем тело ответа
	log.Println(resp.Body)
	var versionList []statusResponse
	err = json.NewDecoder(resp.Body).Decode(&versionList)
	if err != nil {
		return nil, err
	}

	// возвращаем ответ
	var response []model.DataVersionRemote
	for _, s := range versionList {
		response = append(response, model.DataVersionRemote{
			ID:      s.ID,
			Version: s.Version,
		})
	}

	return response, nil
}
