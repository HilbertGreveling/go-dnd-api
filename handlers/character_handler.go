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

func (h *CharacterHandler) CreateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var character models.Character
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		h.response.WriteError(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	if err := h.repo.Create(&character); err != nil {
		h.response.WriteError(w, "Error creating character", http.StatusInternalServerError)
		return
	}

	h.response.WriteJSON(w, character, "OK", http.StatusOK)
}

func (h *CharacterHandler) GetAllCharactersHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := h.repo.GetAll()
	if err != nil {
		h.response.WriteError(w, "Unable to retrieve characters", http.StatusInternalServerError)
		return
	}

	h.response.WriteJSON(w, characters, "OK", http.StatusOK)
}

func (h *CharacterHandler) GetCharacterHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.WriteError(w, "Invalid character ID", http.StatusInternalServerError)
		return
	}

	character, err := h.repo.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving character", http.StatusInternalServerError)
		return
	}

	if character == nil {
		h.response.WriteError(w, "Character not found", http.StatusInternalServerError)
		return
	}

	h.response.WriteJSON(w, character, "OK", http.StatusOK)
}

func (h *CharacterHandler) UpdateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.WriteError(w, "Invalid character ID", http.StatusInternalServerError)
		return
	}

	var updatedCharacter models.Character
	if err := json.NewDecoder(r.Body).Decode(&updatedCharacter); err != nil {
		h.response.WriteError(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	existingCharacter, err := h.repo.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving character", http.StatusInternalServerError)
		return
	}

	updatedCharacter.ID = existingCharacter.ID

	if err := h.repo.Update(&updatedCharacter); err != nil {
		h.response.WriteError(w, "Error updating character", http.StatusInternalServerError)
		return
	}

	h.response.WriteJSON(w, updatedCharacter, "Ok", http.StatusOK)
}

func (h *CharacterHandler) DeleteCharacterHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.WriteError(w, "Invalid character ID", http.StatusInternalServerError)
		return
	}

	_, err = h.repo.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving character", http.StatusInternalServerError)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		h.response.WriteError(w, "Error deleting character", http.StatusInternalServerError)
		return
	}

	h.response.WriteJSON(w, nil, "Character deleted", http.StatusOK)
}
