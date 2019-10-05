package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ViewLogin(w http.ResponseWriter, r *http.Request) {
	message := &Basic{"[GET] Login page here"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}

func ViewRegister(w http.ResponseWriter, r *http.Request) {
	message := &Basic{"[GET] Registration page here"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}
