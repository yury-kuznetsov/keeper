package model

import "github.com/google/uuid"

// User is a struct type that represents a piece of data
// with an ID, Email and a Password.
type User struct {
	ID       uuid.UUID
	Email    string
	Password string
}
