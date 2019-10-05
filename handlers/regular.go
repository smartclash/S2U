package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Basic struct {
	Message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	message := &Basic{Message: "Hello World"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}

func Privacy(w http.ResponseWriter, r *http.Request) {
	message := &Basic{"Privacy policy page here"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}
