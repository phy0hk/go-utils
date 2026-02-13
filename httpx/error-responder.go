package httpx

import (
	"encoding/json"
	"net/http"
)

func ErrorResponder(w http.ResponseWriter, r *http.Request, error_message string, responseCode int) {
	response := ResponseMessage{
		Type:    Error,
		Message: error_message,
	}
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(response)
}
