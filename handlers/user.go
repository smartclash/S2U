package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func UserDashboard(w http.ResponseWriter, r *http.Request) {
	message := &Basic{Message: "[GET] User dashboard here"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}

func UserProfile(w http.ResponseWriter, r *http.Request) {
	message := &Basic{Message: "[GET] User profile page here"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}
