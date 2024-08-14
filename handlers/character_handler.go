package handlers

import (
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/repository"
	"github.com/hilbertgreveling/dnd-character-api/responses"
)

type CharacterHandler struct {
	repo     repository.CharacterRepository
	response responses.JSONResponse
}

func NewCharacterHandler(repo repository.CharacterRepository, response responses.JSONResponse) *CharacterHandler {
	return &CharacterHandler{
		repo:     repo,
		response: response,
	}
}

func (h *CharacterHandler) GetAllCharactersHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := h.repo.GetAll()
	if err != nil {
		h.response.Error(w, "Unable to retrieve characters", http.StatusInternalServerError)
		return
	}

	h.response.Send(w, characters, http.StatusOK)
}
