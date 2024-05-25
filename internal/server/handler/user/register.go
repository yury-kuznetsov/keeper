package user

import (
	"encoding/json"
	"gophkeeper/internal/server/middleware"
	"net/http"
)

type registerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// RegisterHandler handles the HTTP POST request for user registration. It accepts the request, decodes it,
// registers the user, generates a token, saves it in a cookie and adds it to the Authorization header
// of the response. Finally, it sets the HTTP status code to 200 OK.
//
// The function expects two parameters:
// - svc: an object that implements the Service interface, which provides user registration and login functionality.
// - jwt: an object that implements the JWTService interface, which provides functionality for generating JWT tokens.
//
// The function returns an http.HandlerFunc that handles the registration request.
//
// Usage example:
//
//	r.Post("/api/user/register", userHandler.RegisterHandler(userSvc, jwtSvc))
func RegisterHandler(svc Service, jwt JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// принимаем запрос
		var req registerRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// регистрируем пользователя
		userID, err := svc.Register(r.Context(), req.Email, req.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// генерируем токен и сохраняем в куки
		token := jwt.GenerateToken(userID)
		http.SetCookie(w, &http.Cookie{
			Name:  middleware.CookieKey,
			Value: token,
			Path:  "/",
		})
		w.Header().Set("Authorization", token)

		w.WriteHeader(http.StatusOK)
	}
}
