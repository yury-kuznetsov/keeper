package api

import (
	"encoding/json"
	"errors"
	"gophkeeper/internal/client/model"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type pullResponse struct {
	ID       uuid.UUID `json:"id"`
	Category int       `json:"category"`
	Data     []byte    `json:"data"`
	Version  int       `json:"version"`
}

// Pull sends a request to the server to retrieve data with the specified ID.
// It prepares the request, sends it, and handles possible errors.
// If the request is successful, it decrypts the received data and returns it as a model.Data.
// The Pull method returns an error if the request fails.
func (c *Client) Pull(id uuid.UUID) (model.Data, error) {
	path := "/api/data/pull?id=" + id.String()

	// создаем новый запрос
	req, err := http.NewRequest("GET", c.address+path, nil)
	if err != nil {
		return model.Data{}, err
	}

	// добавляем заголовок с JWT-токеном
	req.Header.Add("Authorization", "Bearer "+loadToken())

	// отправляем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.Data{}, err
	}
	defer resp.Body.Close()

	// извлекаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.Data{}, err
	}

	// проверяем возможные ошибки
	if resp.StatusCode != http.StatusOK {
		return model.Data{}, errors.New(string(body))
	}

	// обрабатываем ответ
	var response pullResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return model.Data{}, err
	}

	// дешифруем данные перед возвратом
	dataDecrypt, err := decrypt(string(response.Data))
	if err != nil {
		return model.Data{}, err
	}

	return model.Data{
		ID:       response.ID,
		Category: response.Category,
		Data:     []byte(dataDecrypt),
		Version:  response.Version,
	}, nil
}
