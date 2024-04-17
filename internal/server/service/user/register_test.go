package user

import (
	"context"
	"errors"
	mock_user "gophkeeper/internal/server/mock/user"
	"gophkeeper/internal/server/model"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/google/uuid"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	m := mock_user.NewMockRepository(ctrl)
	m.EXPECT().FindByEmail(ctx, "user@already.exists").Return(model.User{ID: uuid.New()}, nil)
	m.EXPECT().FindByEmail(ctx, "user1@not.exists").Return(model.User{}, nil)
	m.EXPECT().FindByEmail(ctx, "user2@not.exists").Return(model.User{}, nil)
	m.EXPECT().Create(ctx, "user1@not.exists", gomock.Any()).Return(uuid.Nil, errors.New("failure created"))
	m.EXPECT().Create(ctx, "user2@not.exists", gomock.Any()).Return(uuid.New(), nil)

	svc := NewService(m)

	tests := []struct {
		name     string
		email    string
		password string
		error    string
	}{
		{
			name:     "Incorrect email",
			email:    "not-email",
			password: "password",
			error:    "incorrect email",
		},
		{
			name:     "Short password",
			email:    "my@email.com",
			password: "short",
			error:    "password is too short",
		},
		{
			name:     "Long password",
			email:    "my@email.com",
			password: strings.Repeat("long", 72),
			error:    "password is too long",
		},
		{
			name:     "User already exists",
			email:    "user@already.exists",
			password: "password",
			error:    "user already exists",
		},
		{
			name:     "Failure created",
			email:    "user1@not.exists",
			password: "password",
			error:    "failure created",
		},
		{
			name:     "Success created",
			email:    "user2@not.exists",
			password: "password",
			error:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := svc.Register(context.Background(), tt.email, tt.password)
			if tt.error != "" && tt.error != err.Error() {
				t.Error("expected an error but did not get one")
			}
			if tt.error == "" && err != nil {
				t.Errorf("did not expect an error but got one: %v", err)
			}
		})
	}
}
