package responses

import "net/http"

type Response interface {
	WriteResponse(w http.ResponseWriter, data interface{}, message string, statusCode int)
	WriteError(w http.ResponseWriter, message string, statusCode int)
}
