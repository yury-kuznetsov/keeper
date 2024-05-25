package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// JWTService is an interface that defines methods for interacting with a JWT system.
type JWTService interface {
	GetUserID(token string) uuid.UUID
}

type contextKey string

const (

	// CookieKey is a constant that represents the key name used for the cookie in
	// the HTTP request and response headers.
	CookieKey = "token"

	// keyUserID is a constant that represents the context key used for storing the user ID.
	keyUserID contextKey = "userID"
)

// AuthMiddleware is a middleware function that performs authentication and authorization.
// It extracts the token from the request header or cookie, validates it using the provided JWTService,
// and sets the user ID in the request context for further request handling.
// If the token is missing or invalid, it returns an HTTP 401 Unauthorized error.
// Example usage:
// r.Use(AuthMiddleware(jwtSvc))
func AuthMiddleware(jwtSvc JWTService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// извлекаем токен из заголовка или куки
			token := findToken(r)
			if token == "" {
				http.Error(w, "token required", http.StatusUnauthorized)
				return
			}

			// извлекаем идентификатор пользователя
			userID := jwtSvc.GetUserID(token)
			if userID == uuid.Nil {
				http.Error(w, "authorization required", http.StatusUnauthorized)
				return
			}

			// передаем в контекст для обработчиков
			ctx := context.WithValue(r.Context(), keyUserID, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func findToken(r *http.Request) string {
	// ищем в заголовке
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if tokenString != "" {
		return tokenString
	}

	// ищем в куки
	cookie, err := r.Cookie(CookieKey)
	if err != nil {
		return ""
	}

	return cookie.Value
}

// GetUserID extracts the user ID from the context. If the user ID is not found or
// is not of type uuid.UUID, it returns uuid.Nil.
func GetUserID(ctx context.Context) uuid.UUID {
	id, ok := ctx.Value(keyUserID).(uuid.UUID)
	if !ok {
		return uuid.Nil
	}

	return id
}
