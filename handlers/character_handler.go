package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/models"
)

type CharacterHandler struct{}

func (h *CharacterHandler) GetAllCharactersHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := models.GetAllCharacters()
	if err != nil {
		http.Error(w, "Failed to retrieve characters", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(characters)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
