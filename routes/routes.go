package routes

import (
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/handlers"
	"github.com/hilbertgreveling/dnd-character-api/repository"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Ping
	pingHandler := handlers.NewPingHandler()

	mux.HandleFunc("GET /ping", pingHandler.Ping)

	// Character

	repo := repository.NewCharacterRepositorySQLite()
	characterHandler := handlers.NewCharacterHandler(repo)

	mux.HandleFunc("GET /characters", characterHandler.GetAllCharactersHandler)

	return mux
}
