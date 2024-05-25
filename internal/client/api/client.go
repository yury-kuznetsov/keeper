package api

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

// Client структура для отправки запросов на сервер
type Client struct {
	address string
}

// NewClient creates a new instance of the Client struct with the given address.
// The Client struct is used for sending requests to the server.
// The address parameter specifies the server address to which the requests will be sent.
// The returned value is a pointer to the created Client instance.
func NewClient(address string) *Client {
	return &Client{address: address}
}

func saveToken(token string) error {
	return os.WriteFile("token.txt", []byte(token), 0600)
}

func loadToken() string {
	bytes, err := os.ReadFile("token.txt")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func getSecretKey() string {
	// в реальности у пользователя будет где-то храниться ключ шифрования,
	// который он будет переносить с устройства на устроство
	bytes, err := os.ReadFile("secret.txt")
	if err != nil {
		return ""
	}
	return string(bytes)
}

func encrypt(data []byte) (ciphertext string, err error) {
	secretKey := getSecretKey()

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	cypher := gcm.Seal(nonce, nonce, data, nil)
	ciphertext = hex.EncodeToString(cypher)

	return
}

func decrypt(cipherHex string) (plaintext string, err error) {
	secretKey := getSecretKey()

	ciphertext, err := hex.DecodeString(cipherHex)
	if err != nil {
		return
	}

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		err = errors.New("ciphertext too short")
		return
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	bytes, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return
	}

	plaintext = string(bytes)

	return
}
