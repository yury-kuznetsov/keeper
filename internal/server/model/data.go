package model

import "github.com/google/uuid"

// Data is a struct type that represents a piece of data
// with an ID, UserID, Category, the actual Data, and a Version number.
type Data struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Category int
	Data     []byte
	Version  int
}

// DataVersion is a struct type that represents a piece of data
// with an ID and a Version number.
type DataVersion struct {
	ID      uuid.UUID
	Version int
}
