package responses

import (
	"encoding/json"
	"net/http"
)

type JSONResponse interface {
	Send(w http.ResponseWriter, data interface{}, message string, statusCode int)
	Error(w http.ResponseWriter, message string, statusCode int)
}

type DefaultJSONResponse struct{}

func NewDefaultJSONResponse() *DefaultJSONResponse {
	return &DefaultJSONResponse{}
}

func (r *DefaultJSONResponse) Send(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"data":    data,
		"message": message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func (r *DefaultJSONResponse) Error(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errResponse := map[string]string{"error": message}
	if err := json.NewEncoder(w).Encode(errResponse); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}
