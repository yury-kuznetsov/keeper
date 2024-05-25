package data

import (
	"context"
	mock_data "gophkeeper/internal/server/mock/data"
	"gophkeeper/internal/server/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ID := uuid.New()
	userID := uuid.New()

	// хранящаяся в базе запись
	data := model.Data{
		ID:       ID,
		UserID:   userID,
		Category: 0,
		Data:     []byte("DATA"),
		Version:  0,
	}

	m := mock_data.NewMockRepository(ctrl)
	m.EXPECT().Find(context.Background(), ID, userID).Return(data, nil).Times(3)
	m.EXPECT().Save(context.Background(), gomock.Any()).Return(nil)

	svc := NewService(m)

	tests := []struct {
		name      string
		id        uuid.UUID
		userID    uuid.UUID
		data      []byte
		version   int
		isUpdated bool
		error     string
	}{
		{
			name:      "Incorrect version",
			id:        ID,
			userID:    userID,
			data:      []byte("NEW DATA"),
			version:   1, // версия входящего сообщения не совпадает с серверной
			isUpdated: false,
			error:     "incorrect version",
		},
		{
			name:      "Not updated",
			id:        ID,
			userID:    userID,
			data:      []byte("DATA"),
			version:   0,
			isUpdated: false, // обновления не будет, ведь данные совпадают
			error:     "",
		},
		{
			name:      "Updated",
			id:        ID,
			userID:    userID,
			data:      []byte("NEW DATA"),
			version:   0,
			isUpdated: true, // версия данных должна увеличиться
			error:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := svc.Push(context.Background(), tt.id, tt.userID, 0, tt.data, tt.version)
			assert.Equal(t, v != tt.version, tt.isUpdated)
			if tt.error != "" {
				assert.Equal(t, err.Error(), tt.error)
			}
		})
	}
}
