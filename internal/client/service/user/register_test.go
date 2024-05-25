package user

import (
	mock_user "gophkeeper/internal/client/mock/user"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_user.NewMockClient(ctrl)
	m.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)

	svc := NewService(m)

	tests := []struct {
		name     string
		email    string
		password string
		error    string
	}{
		{
			name:     "Incorrect email",
			email:    "incorrect-email",
			password: "password",
			error:    "incorrect email",
		},
		{
			name:     "Short password",
			email:    "user@email.com",
			password: "short",
			error:    "password is too short",
		},
		{
			name:     "Long password",
			email:    "user@email.com",
			password: strings.Repeat("long", 72),
			error:    "password is too long",
		},
		{
			name:     "Success",
			email:    "user@email.com",
			password: "password",
			error:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Register(tt.email, tt.password)
			if tt.error == "" {
				assert.Nil(t, err)
			} else {
				assert.Equal(t, err.Error(), tt.error)
			}
		})
	}
}
