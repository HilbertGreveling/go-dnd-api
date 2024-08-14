package routes

import (
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Ping
	pingHandler := handlers.NewPingHandler()

	mux.HandleFunc("GET /ping", pingHandler.Ping)

	// Character

	characterHandler := &handlers.CharacterHandler{}

	mux.HandleFunc("GET /characters", characterHandler.GetAllCharactersHandler)

	return mux
}
