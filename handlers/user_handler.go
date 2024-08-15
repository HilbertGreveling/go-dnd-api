package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/repository"
	"github.com/hilbertgreveling/dnd-character-api/responses"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo     repository.UserRepository
	response responses.JSONResponse
}

func NewUserHandler(repo repository.UserRepository, response responses.JSONResponse) *UserHandler {
	return &UserHandler{
		repo:     repo,
		response: response,
	}
}

func (h *UserHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.response.WriteError(w, "Invalid request payload", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		h.response.WriteError(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := h.repo.Create(&user); err != nil {
		h.response.WriteError(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	userResponse := models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	h.response.WriteJSON(w, userResponse, "User registered successfully", http.StatusCreated)
}

func (h *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.response.WriteError(w, "Invalid User ID", http.StatusInternalServerError)
		return
	}

	user, err := h.repo.GetByID(id)
	if err != nil {
		h.response.WriteError(w, "Error retrieving user", http.StatusInternalServerError)
		return
	}

	if user == nil {
		h.response.WriteError(w, "User not found", http.StatusInternalServerError)
		return
	}

	h.response.WriteJSON(w, user, "OK", http.StatusOK)
}
