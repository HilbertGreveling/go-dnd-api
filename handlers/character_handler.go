package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/repository"
)

type CharacterHandler struct {
	repo repository.CharacterRepository
}

func NewCharacterHandler(repo repository.CharacterRepository) *CharacterHandler {
	return &CharacterHandler{repo: repo}
}

func (h *CharacterHandler) GetAllCharactersHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, "Unable to retrieve characters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(characters)
}
