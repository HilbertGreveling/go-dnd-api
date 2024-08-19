package handlers

import (
	"net/http"
	"strconv"

	"github.com/hilbertgreveling/dnd-character-api/responses"
	"github.com/hilbertgreveling/dnd-character-api/services"
)

type UserHandler struct {
	service  services.UserService
	response responses.Response
}

func NewUserHandler(service services.UserService, response responses.Response) *UserHandler {
	return &UserHandler{
		service:  service,
		response: response,
	}
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.WriteError(w, "Invalid User ID", http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		h.response.WriteError(w, "User not found", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, user, "OK", http.StatusOK)
}
