package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/config"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg := config.LoadConfig()

		if r.Header.Get("Authorization") != cfg.SecretKey {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(ErrorResponse{Message: "Permission Denied"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
