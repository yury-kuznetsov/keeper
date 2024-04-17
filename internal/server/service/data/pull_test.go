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

func TestPull(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()
	userID := uuid.New()

	m := mock_data.NewMockRepository(ctrl)
	m.EXPECT().Find(context.Background(), id, userID).Return(model.Data{ID: id}, nil)

	svc := NewService(m)

	t.Run("Pull", func(t *testing.T) {
		data, err := svc.Pull(context.Background(), id, userID)
		assert.Equal(t, data.ID, id)
		assert.Nil(t, err, "Error is not nil")
	})
}
