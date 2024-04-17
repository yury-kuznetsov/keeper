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

func TestPull(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idNotExists := uuid.New()
	idExists := uuid.New()

	r := mock_data.NewMockRepository(ctrl)
	c := mock_data.NewMockClient(ctrl)

	r.EXPECT().Save(gomock.Any()).Return(nil)

	c.EXPECT().Pull(idNotExists).Return(model.Data{}, errors.New("REMOTE_ERROR"))
	c.EXPECT().Pull(idExists).Return(model.Data{
		ID:       idExists,
		Category: model.CategoryBinary,
		Data:     []byte("DATA"),
		Version:  1,
	}, nil)

	svc := NewService(r, c)

	tests := []struct {
		name  string
		id    uuid.UUID
		error string
	}{
		{
			name:  "Remote error",
			id:    idNotExists,
			error: "REMOTE_ERROR",
		},
		{
			name:  "Update data",
			id:    idExists,
			error: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Pull(tt.id)
			if tt.error == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, err.Error(), tt.error)
			}
		})
	}
}
