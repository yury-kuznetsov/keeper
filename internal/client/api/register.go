package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register sends a request to the server to register a new user using the provided email and password.
// It prepares the request, sends it, and handles possible errors.
// If the registration is successful, it saves the JWT token received in the response header.
// The Register method returns an error if the registration fails.
func (c *Client) Register(email, password string) error {
	path := "/api/user/register"

	// подготавливаем запрос
	request := registerRequest{
		Email:    email,
		Password: password,
	}
	jsonData, _ := json.Marshal(request)

	// отправляем запрос
	resp, err := http.Post(c.address+path, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// проверяем возможные ошибки
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("registration failed: %s", string(body))
	}

	// сохраняем JWT
	err = saveToken(resp.Header.Get("Authorization"))
	if err != nil {
		return err
	}

	return nil
}
