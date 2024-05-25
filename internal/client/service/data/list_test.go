package data

import (
	mock_data "gophkeeper/internal/client/mock/data"
	"gophkeeper/internal/client/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	r := mock_data.NewMockRepository(ctrl)
	c := mock_data.NewMockClient(ctrl)

	models := []model.Data{
		{
			ID:       uuid.New(),
			Category: model.CategoryCard,
			Data:     []byte("BANK_CARD"),
			Version:  0,
		},
	}

	r.EXPECT().FindAll().Return(models, nil)

	svc := NewService(r, c)

	t.Run("List", func(t *testing.T) {
		data, err := svc.List()
		assert.Equal(t, data, models)
		assert.Nil(t, err)
	})
}
