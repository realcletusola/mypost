package handlers

import (
	"encoding/json"
	"net/http"
)

// Create function to decode Json
func decodeJSONBody(w http.ResponseWriter, r *http.Request, dest interface{}) error {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	return decoder.Decode(dest)
}

// Create function to encode Json
func encodeJSONBody(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}