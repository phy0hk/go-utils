package httpx

import (
	"net/http"
	"slices"
)

func CheckAllowedMethods(w http.ResponseWriter, r *http.Request, allowedMethods []string) bool {
	if slices.Contains(allowedMethods, r.Method) {
		return true
	}
	ErrorResponder(w, r, "Invalid Method", http.StatusMethodNotAllowed)
	return false
}
