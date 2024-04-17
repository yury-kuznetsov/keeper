package data

import (
	"errors"
	mock_data "gophkeeper/internal/client/mock/data"
	"gophkeeper/internal/client/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idNotExists := uuid.New()
	idExists := uuid.New()

	r := mock_data.NewMockRepository(ctrl)
	c := mock_data.NewMockClient(ctrl)

	r.EXPECT().Find(idNotExists).Return(model.Data{})
	r.EXPECT().Find(idExists).Return(model.Data{ID: idExists, Version: 1}).Times(3)
	r.EXPECT().Save(gomock.Any()).Return(nil)

	c.EXPECT().Push(gomock.Any()).Return(0, errors.New("PUSH_ERROR"))
	c.EXPECT().Push(gomock.Any()).Return(1, nil)
	c.EXPECT().Push(gomock.Any()).Return(2, nil)

	svc := NewService(r, c)

	tests := []struct {
		name  string
		id    uuid.UUID
		error string
	}{
		{
			name:  "Not found",
			id:    idNotExists,
			error: "entity not found",
		},
		{
			name:  "Push error",
			id:    idExists,
			error: "PUSH_ERROR",
		},
		{
			name:  "Equal versions",
			id:    idExists,
			error: "",
		},
		{
			name:  "Up version",
			id:    idExists,
			error: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Push(tt.id)
			if tt.error == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, err.Error(), tt.error)
			}
		})
	}
}
