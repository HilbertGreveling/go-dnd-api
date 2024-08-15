package routes

import (
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/db"
	"github.com/hilbertgreveling/dnd-character-api/handlers"
	"github.com/hilbertgreveling/dnd-character-api/repository"
	"github.com/hilbertgreveling/dnd-character-api/responses"
)

func SetupRoutes(mux *http.ServeMux) *http.ServeMux {
	// Ping
	pingHandler := handlers.NewPingHandler()

	mux.HandleFunc("GET /ping", pingHandler.Ping)

	db := db.GetDB()
	jsonResponse := responses.NewDefaultJSONResponse()

	// Character
	repo := repository.NewCharacterRepositorySQLite(db)
	characterHandler := handlers.NewCharacterHandler(repo, jsonResponse)

	mux.HandleFunc("POST /characters/new", characterHandler.CreateCharacterHandler)
	mux.HandleFunc("GET /characters", characterHandler.GetAllCharactersHandler)
	mux.HandleFunc("GET /characters/{id}", characterHandler.GetCharacterHandler)
	mux.HandleFunc("PUT /characters/{id}", characterHandler.UpdateCharacterHandler)
	mux.HandleFunc("DELETE /characters/{id}", characterHandler.DeleteCharacterHandler)

	return mux
}
