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

	// Ping
	pingHandler := NewPingHandler()

	mux.Handle("GET /ping", middleware.AuthMiddleware(http.HandlerFunc(pingHandler.Ping)))

	services := services.SetupServices()

	// User
	userHandler := NewUserHandler(services.UserService, response)

	mux.HandleFunc("POST /register", userHandler.RegisterUserHandler)
	mux.HandleFunc("POST /login", userHandler.LoginHandler)

	mux.HandleFunc("GET /users/{id}", userHandler.GetUserHandler)

	// Character
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
