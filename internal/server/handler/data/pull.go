package data

import (
	"encoding/json"
	"gophkeeper/internal/server/middleware"
	"net/http"

	"github.com/google/uuid"
)

type pullResponse struct {
	ID       uuid.UUID `json:"id"`
	Category int       `json:"category"`
	Data     []byte    `json:"data"`
	Version  int       `json:"version"`
}

// PullHandler is a function that handles the HTTP POST request for pulling data.
// It decodes the request body into a pullRequest struct, extracts the user ID from the context,
// and calls the Pull method of the provided Service interface to retrieve the data.
// It then encodes the data into a pullResponse struct and writes it to the response as JSON.
// If any error occurs during the process, it returns an appropriate HTTP error response.
//
// It takes an instance of the Service interface as a parameter and returns an http.HandlerFunc.
// The returned http.HandlerFunc serves as the actual handler for the HTTP request.
func PullHandler(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r.Context())

		// принимаем запрос
		id, err := uuid.Parse(r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// извлекаем данные пользователя
		data, err := svc.Pull(r.Context(), id, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// возвращаем ответ
		jsonData, err := json.Marshal(pullResponse{
			ID:       data.ID,
			Category: data.Category,
			Data:     data.Data,
			Version:  data.Version,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
