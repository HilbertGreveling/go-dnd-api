package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/hilbertgreveling/dnd-character-api/models"
	"github.com/hilbertgreveling/dnd-character-api/responses"
	"github.com/hilbertgreveling/dnd-character-api/services"
)

type AuthHandler struct {
	service  services.AuthService
	response responses.Response
}

func NewAuthHandler(service services.AuthService, response responses.Response) *AuthHandler {
	return &AuthHandler{
		service:  service,
		response: response,
	}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var registrationData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&registrationData); err != nil {
		h.response.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user := models.User{
		Username: registrationData.Username,
	}

	err := h.service.RegisterUser(user, registrationData.Password)
	if err != nil {
		h.response.WriteError(w, "User creation failed", http.StatusInternalServerError)
		return
	}

	h.response.WriteResponse(w, user, "User registered successfully", http.StatusOK)
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {

	//
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		h.response.WriteError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.service.LoginUser(loginData.Username, loginData.Password)
	if err != nil {
		h.response.WriteError(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	h.response.WriteResponse(w, map[string]string{"token": token}, "Login successful", http.StatusOK)
}
