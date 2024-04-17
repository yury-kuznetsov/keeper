package data

import (
	"encoding/json"
	"gophkeeper/internal/server/middleware"
	"net/http"

	"github.com/google/uuid"
)

type statusResponse struct {
	ID      uuid.UUID `json:"id"`
	Version int       `json:"version"`
}

// StatusHandler is a function that handles the HTTP request for retrieving the status
// of a user's data. It expects a Service interface as a parameter. It retrieves the user ID
// from the request context using the middleware.GetUserID function. It calls the Status method
// of the Service interface to get the data status. It prepares the response by converting
// the data status into a slice of statusResponse structs. It converts the response to JSON
// and writes it to the response writer with the appropriate content type header.
func StatusHandler(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := middleware.GetUserID(r.Context())

		// извлекаем данные пользователя
		data, err := svc.Status(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// подготавливаем ответ
		var response []statusResponse
		for _, d := range data {
			response = append(response, statusResponse{
				ID:      d.ID,
				Version: d.Version,
			})
		}

		// возвращаем ответ
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}
