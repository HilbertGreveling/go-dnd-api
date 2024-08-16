package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/responses"
	"github.com/hilbertgreveling/dnd-character-api/services"
)

type CharacterHandler struct {
	service  services.CharacterService
	response responses.Response
}

func NewCharacterHandler(service services.CharacterService, response responses.Response) *CharacterHandler {
	return &CharacterHandler{
		service:  service,
		response: response,
	}
}

func (h *CharacterHandler) CreateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var character models.Character
	if err := json.NewDecoder(r.Body).Decode(&character); err != nil {
		h.response.WriteError(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	if err := h.service.Create(&character); err != nil {
		h.response.WriteError(w, "Error creating character", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, character, "OK", http.StatusOK)
}

func (h *CharacterHandler) GetAllCharactersHandler(w http.ResponseWriter, r *http.Request) {
	characters, err := h.service.GetAll()
	if err != nil {
		h.response.WriteError(w, "Unable to retrieve characters", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, characters, "OK", http.StatusOK)
}

func (h *CharacterHandler) GetCharacterHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.WriteError(w, "Invalid character ID", http.StatusInternalServerError)
		return
	}

	character, err := h.service.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving character", http.StatusInternalServerError)
		return
	}

	if character == nil {
		h.response.WriteError(w, "Character not found", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, character, "OK", http.StatusOK)
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

	existingCharacter, err := h.service.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving character", http.StatusInternalServerError)
		return
	}

	updatedCharacter.ID = existingCharacter.ID

	if err := h.service.Update(&updatedCharacter); err != nil {
		h.response.WriteError(w, "Error updating character", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, updatedCharacter, "Ok", http.StatusOK)
}

func (h *CharacterHandler) DeleteCharacterHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.WriteError(w, "Invalid character ID", http.StatusInternalServerError)
		return
	}

	_, err = h.service.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving character", http.StatusInternalServerError)
		return
	}

	if err := h.service.Delete(id); err != nil {
		h.response.WriteError(w, "Error deleting character", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, nil, "Character deleted", http.StatusOK)
}
