package webtoken

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

const (
	// SecretKey is a constant string representing the secret key used for JWT token signing and verification.
	SecretKey = "SECRET_KEY"

	// Duration represents the duration used for JWT token expiry.
	Duration = time.Hour
)

// JWTService is a type that represents a JWT service
type JWTService struct{}

// Claims represents the custom claims for a JWT token, which includes
// the standard RegisteredClaims and an additional UserID field.
type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID
}

// NewJWTService returns a new instance of the JWTService struct.
// It initializes the JWTService struct with an empty state and returns a pointer to it.
func NewJWTService() *JWTService {
	return &JWTService{}
}

// GenerateToken generates a new JWT token for the given user ID.
// It checks if the user ID is not nil and returns an empty string if it is.
// The token is created with the HS256 signing method and includes the user ID and expiry time.
// The token is signed using the secret key and returned as a string.
func (s *JWTService) GenerateToken(userID uuid.UUID) string {
	if userID == uuid.Nil {
		return ""
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(Duration)),
		},
		UserID: userID,
	})

	tokenString, _ := token.SignedString([]byte(SecretKey))

	return tokenString
}

// GetUserID parses the provided JWT token string and returns the user ID contained in the token.
// It uses the HS256 signing method and the provided secret key to verify the token's authenticity.
// If the token is invalid or an error occurs during parsing, it returns the zero value of the UUID type.
// Otherwise, it returns the user ID extracted from the token's claims.
func (s *JWTService) GetUserID(tokenString string) uuid.UUID {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return uuid.Nil
	}

	if !token.Valid {
		return uuid.Nil
	}

	return claims.UserID
}
