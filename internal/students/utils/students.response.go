package response_utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON sends a JSON response.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}
