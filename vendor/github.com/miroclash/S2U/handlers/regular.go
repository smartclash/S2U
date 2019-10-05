package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type basic struct {
	Message string `json:"message"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	message := &basic{Message: "Hello World"}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}
