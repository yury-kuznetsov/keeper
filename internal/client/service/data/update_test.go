package data

import (
	mock_data "gophkeeper/internal/client/mock/data"
	"gophkeeper/internal/client/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	idNotExists := uuid.New()
	idExists := uuid.New()

	r := mock_data.NewMockRepository(ctrl)
	c := mock_data.NewMockClient(ctrl)

	r.EXPECT().Find(idNotExists).Return(model.Data{})
	r.EXPECT().Find(idExists).Return(model.Data{ID: idExists, Data: []byte("DATA")}).Times(2)
	r.EXPECT().Save(gomock.Any()).Return(nil)

	svc := NewService(r, c)

	tests := []struct {
		name  string
		id    uuid.UUID
		data  []byte
		error string
	}{
		{
			name:  "Not found",
			id:    idNotExists,
			data:  []byte("NEW DATA"),
			error: "entity not found",
		},
		{
			name:  "Equal data",
			id:    idExists,
			data:  []byte("DATA"),
			error: "",
		},
		{
			name:  "Update data",
			id:    idExists,
			data:  []byte("NEW DATA"),
			error: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Update(tt.id, tt.data)
			if tt.error == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, err.Error(), tt.error)
			}
		})
	}
}
