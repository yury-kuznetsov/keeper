package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login sends a request to the server to authenticate a user using email and password.
// It prepares the request, sends it, and handles possible errors.
// If the login is successful, it saves the JWT token received in the response header.
// The login function returns an error if the login fails.
func (c *Client) Login(email, password string) error {
	path := "/api/user/login"

	// подготавливаем запрос
	request := loginRequest{
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
		return fmt.Errorf("login failed: %s", string(body))
	}

	// сохраняем JWT
	err = saveToken(resp.Header.Get("Authorization"))
	if err != nil {
		return err
	}

	return nil
}
