package webtoken

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func TestJWTService_GenerateToken(t *testing.T) {
	tests := []struct {
		name      string
		userID    uuid.UUID
		wantEmpty bool
	}{
		{
			name:      "ValidUserID",
			userID:    uuid.New(),
			wantEmpty: false,
		},
		{
			name:      "EmptyUserID",
			userID:    uuid.Nil,
			wantEmpty: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serv := NewJWTService()
			got := serv.GenerateToken(test.userID)
			if test.wantEmpty && got != "" {
				t.Fatalf("Expected empty token but got %q", got)
			}
			if !test.wantEmpty && got == "" {
				t.Fatal("Expected a token but got an empty string")
			}
		})
	}
}

func TestJWTService_GetUserID(t *testing.T) {
	userID := uuid.New()
	tests := []struct {
		name     string
		tokenStr string
		want     uuid.UUID
	}{
		{
			name:     "ValidToken",
			tokenStr: newToken(userID),
			want:     userID,
		},
		{
			name:     "InvalidToken",
			tokenStr: "",
			want:     uuid.Nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			serv := NewJWTService()
			got := serv.GetUserID(test.tokenStr)
			if got != test.want {
				t.Fatalf("Expected %q but got %q", test.want, got)
			}
		})
	}
}

func newToken(userID uuid.UUID) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(Duration)),
		},
		UserID: userID,
	})
	tokenStr, _ := token.SignedString([]byte(SecretKey))
	return tokenStr
}
