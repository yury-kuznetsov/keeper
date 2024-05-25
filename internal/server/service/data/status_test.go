package data

import (
	"context"
	mock_data "gophkeeper/internal/server/mock/data"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userID := uuid.New()

	m := mock_data.NewMockRepository(ctrl)
	m.EXPECT().Versions(context.Background(), userID).Return(nil, nil)

	svc := NewService(m)

	t.Run("Status", func(t *testing.T) {
		data, err := svc.Status(context.Background(), userID)
		assert.Nil(t, data, "Returned data is not nil")
		assert.Nil(t, err, "Error is not nil")
	})
}
