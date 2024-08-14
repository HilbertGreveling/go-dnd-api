package routes

import (
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/handlers"
	"github.com/hilbertgreveling/dnd-character-api/repository"
	"github.com/hilbertgreveling/dnd-character-api/responses"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Ping
	pingHandler := handlers.NewPingHandler()

	mux.HandleFunc("GET /ping", pingHandler.Ping)

	jsonResponse := responses.NewDefaultJSONResponse()

	// Character
	repo := repository.NewCharacterRepositorySQLite()
	characterHandler := handlers.NewCharacterHandler(repo, jsonResponse)

	mux.HandleFunc("GET /characters", characterHandler.GetAllCharactersHandler)

	return mux
}
