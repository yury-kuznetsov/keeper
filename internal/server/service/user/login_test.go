package user

import (
	"context"
	"errors"
	mock_user "gophkeeper/internal/server/mock/user"
	"gophkeeper/internal/server/model"
	"testing"

	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
)

func TestLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user := model.User{
		ID:       uuid.New(),
		Email:    "user@email.com",
		Password: string(passwordHash),
	}

	m := mock_user.NewMockRepository(ctrl)
	m.EXPECT().FindByEmail(ctx, "user@not.found").Return(model.User{}, errors.New("not found"))
	m.EXPECT().FindByEmail(ctx, "user@email.com").Return(user, nil).Times(2)

	svc := NewService(m)

	tests := []struct {
		name     string
		email    string
		password string
		error    string
	}{
		{
			name:     "User not found",
			email:    "user@not.found",
			password: "password",
			error:    "incorrect email or password",
		},
		{
			name:     "Incorrect password",
			email:    "user@email.com",
			password: "incorrect",
			error:    "incorrect email or password",
		},
		{
			name:     "Success login",
			email:    "user@email.com",
			password: "password",
			error:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Login(context.Background(), tt.email, tt.password)
			if tt.error != "" && tt.error != err.Error() {
				t.Error("expected an error but did not get one")
			}
			if tt.error == "" && err != nil {
				t.Errorf("did not expect an error but got one: %v", err)
			}
		})
	}
}
