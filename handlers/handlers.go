package handlers

import (
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/middleware"
	"github.com/hilbertgreveling/dnd-character-api/responses"
	"github.com/hilbertgreveling/dnd-character-api/services"
)

func SetupHandlers() http.Handler {
	mux := http.NewServeMux()
	response := responses.NewDefaultJSONResponse()
	services := services.SetupServices()

	// Ping
	pingHandler := NewPingHandler()

	mux.Handle("GET /ping", middleware.AuthMiddleware(services.UserService, http.HandlerFunc(pingHandler.Ping)))

	// Auth
	authHandler := NewAuthHandler(services.AuthService, response)

	mux.HandleFunc("POST /register", authHandler.RegisterHandler)
	mux.HandleFunc("POST /login", authHandler.LoginHandler)

	// Users
	userHandler := NewUserHandler(services.UserService, response)

	mux.HandleFunc("GET /users/{id}", userHandler.GetUserHandler)

	// Characters
	characterHandler := NewCharacterHandler(services.CharacterService, response)

	mux.HandleFunc("POST /characters/new", characterHandler.CreateCharacterHandler)
	mux.HandleFunc("GET /characters", characterHandler.GetAllCharactersHandler)
	mux.HandleFunc("GET /characters/{id}", characterHandler.GetCharacterHandler)
	mux.HandleFunc("PUT /characters/{id}", characterHandler.UpdateCharacterHandler)
	mux.HandleFunc("DELETE /characters/{id}", characterHandler.DeleteCharacterHandler)

	stack := middleware.CreateStack(
		middleware.CORS,
		middleware.Logging,
	)

	return stack(mux)
}
