package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hilbertgreveling/dnd-character-api/models"
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

func (h *CharacterHandler) GetCharacterHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.Error(w, "Invalid character ID", http.StatusInternalServerError)
		return
	}

	character, err := h.repo.GetByID(id)
	if err != nil {
		h.response.Error(w, "Error retrieving character", http.StatusInternalServerError)
		return
	}

	if character == nil {
		h.response.Error(w, "Character not found", http.StatusInternalServerError)
		return
	}

	h.response.Send(w, character, http.StatusOK)
}

func (h *CharacterHandler) CreateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var character models.Character
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		h.response.Error(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	if err := h.repo.Create(&character); err != nil {
		h.response.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.response.Send(w, character, http.StatusOK)
}
