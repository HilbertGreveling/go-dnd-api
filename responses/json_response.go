package responses

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct{}

func NewDefaultJSONResponse() *JSONResponse {
	return &JSONResponse{}
}

func (r *JSONResponse) WriteResponse(w http.ResponseWriter, data interface{}, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := map[string]interface{}{
		"status":  statusCode,
		"message": message,
		"data":    data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
	}
}

func (r *JSONResponse) WriteError(w http.ResponseWriter, message string, statusCode int) {
	r.WriteResponse(w, nil, message, statusCode)
}
