package routes

import (
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/handlers"
	"github.com/hilbertgreveling/dnd-character-api/responses"
	"github.com/hilbertgreveling/dnd-character-api/services"
)

func SetupRoutes(mux *http.ServeMux) *http.ServeMux {
	// Ping
	pingHandler := handlers.NewPingHandler()

	mux.HandleFunc("GET /ping", pingHandler.Ping)

	jsonResponse := responses.NewDefaultJSONResponse()
	services := services.SetupServices()

	// User
	userHandler := handlers.NewUserHandler(services.UserService, jsonResponse)

	mux.HandleFunc("POST /users/register", userHandler.RegisterUserHandler)
	mux.HandleFunc("GET /users/{id}", userHandler.GetUserHandler)

	// Character
	characterHandler := handlers.NewCharacterHandler(services.CharacterService, jsonResponse)

	mux.HandleFunc("POST /characters/new", characterHandler.CreateCharacterHandler)
	mux.HandleFunc("GET /characters", characterHandler.GetAllCharactersHandler)
	mux.HandleFunc("GET /characters/{id}", characterHandler.GetCharacterHandler)
	mux.HandleFunc("PUT /characters/{id}", characterHandler.UpdateCharacterHandler)
	mux.HandleFunc("DELETE /characters/{id}", characterHandler.DeleteCharacterHandler)

	return mux
}
