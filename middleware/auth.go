package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/hilbertgreveling/dnd-character-api/services"
	"github.com/hilbertgreveling/dnd-character-api/utils"
)

type ContextKey string

const UserIDKey ContextKey = "userID"

func AuthMiddleware(userService services.UserService, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userID := claims.UserID

		user, err := userService.GetByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, user)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
