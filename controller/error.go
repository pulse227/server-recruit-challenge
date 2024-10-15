package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// エラーが発生したときのレスポンス処理をここで行う
func errorHandler(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	slog.ErrorContext(r.Context(), "error occurred", "message", message)

	type ErrorMessage struct {
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(&ErrorMessage{Message: message})
}
