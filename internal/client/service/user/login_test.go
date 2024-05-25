package user

import (
	mock_user "gophkeeper/internal/client/mock/user"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_user.NewMockClient(ctrl)
	m.EXPECT().Login(gomock.Any(), gomock.Any()).Return(nil)

	svc := NewService(m)

	t.Run("Pull", func(t *testing.T) {
		err := svc.Login("user@email.com", "password")
		assert.Nil(t, err)
	})
}
