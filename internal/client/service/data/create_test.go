package data

import (
	mock_data "gophkeeper/internal/client/mock/data"
	"gophkeeper/internal/client/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_data.NewMockRepository(ctrl)
	c := mock_data.NewMockClient(ctrl)

	r.EXPECT().Save(gomock.Any()).Return(nil)

	svc := NewService(r, c)

	t.Run("Create", func(t *testing.T) {
		id, err := svc.Create(model.CategoryCredentials, []byte("LOGIN_AND_PASSWORD"))
		assert.NotEqual(t, id, uuid.Nil)
		assert.Nil(t, err)
	})
}
