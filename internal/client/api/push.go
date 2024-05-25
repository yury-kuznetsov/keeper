package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"gophkeeper/internal/client/model"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type pushRequest struct {
	ID       uuid.UUID `json:"id"`
	Category int       `json:"category"`
	Data     []byte    `json:"data"`
	Version  int       `json:"version"`
}

type pushResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
}

// Push sends a request to the server to push the specified data.
// It prepares the request, encrypts the data, sends it, and handles possible errors.
// If the request is successful, it returns the version of the pushed data.
// The Push method returns an error if the request fails or if encryption fails.
func (c *Client) Push(data model.Data) (int, error) {
	path := "/api/data/push"

	// шифруем данные перед отправкой
	dataEncrypt, err := encrypt(data.Data)
	if err != nil {
		return 0, err
	}

	// подготавливаем запрос
	request := pushRequest{
		ID:       data.ID,
		Category: data.Category,
		Data:     []byte(dataEncrypt),
		Version:  data.Version,
	}
	jsonData, _ := json.Marshal(request)

	// создаем новый запрос
	req, err := http.NewRequest("POST", c.address+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return 0, err
	}

	// добавляем заголовки
	req.Header.Add("Authorization", "Bearer "+loadToken())
	req.Header.Add("Content-Type", "application/json")

	// отправляем запрос
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// извлекаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// проверяем возможные ошибки
	if resp.StatusCode != http.StatusOK {
		return 0, errors.New(string(body))
	}

	// обрабатываем ответ
	var response pushResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	return response.Version, nil
}
