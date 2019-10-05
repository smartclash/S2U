package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ViewCreateLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	message := &Basic{Message: "[GET] Create a new link page will be shown here for user " + vars["id"]}
	if err := json.NewEncoder(w).Encode(message); err != nil {
		_, _ = fmt.Fprint(w, "Unknown error occurred")
	}
}
