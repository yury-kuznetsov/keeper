package data

import (
	"encoding/json"
	"gophkeeper/internal/server/middleware"
	"net/http"

	"github.com/google/uuid"
)

type pushRequest struct {
	ID       uuid.UUID `json:"id"`
	Category int       `json:"category"`
	Data     []byte    `json:"data"`
	Version  int       `json:"version"`
}

type pushResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
}

// PushHandler is http.Handler that handles the push data request.
// It extracts the user ID from the context, decodes the request body,
// saves the user data, and responds with the updated version number.
func PushHandler(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r.Context())

		// принимаем запрос
		var req pushRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		// сохраняем данные пользователя
		version, err := svc.Push(r.Context(), req.ID, userID, req.Category, req.Data, req.Version)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// возвращаем ответ
		res := pushResponse{ID: req.ID, Version: version}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	}
}
