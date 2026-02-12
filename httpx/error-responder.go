package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/phy0hk/go-utils/types"
)

func ErrorResponder(w http.ResponseWriter, r *http.Request, error_message string, responseCode int) {
	response := types.ResponseMessage{
		Type:    types.Error,
		Message: error_message,
	}
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(response)
}
