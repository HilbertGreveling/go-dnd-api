package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/responses"
	"github.com/hilbertgreveling/dnd-character-api/services"
	"github.com/hilbertgreveling/dnd-character-api/utils"
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

func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.response.WriteError(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	if err := h.service.Create(&user); err != nil {
		h.response.WriteError(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, nil, "User registered successfully", http.StatusCreated)
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		h.response.WriteError(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetByUsername(credentials.Username)
	if err != nil {
		h.response.WriteError(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !h.service.CheckPassword(user, credentials.Password) {
		h.response.WriteError(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		h.response.WriteError(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, map[string]string{"token": token}, "Login successful", http.StatusOK)
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
