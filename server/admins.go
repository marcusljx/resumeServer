package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func AddResume(w http.ResponseWriter, r *http.Request) {

	var resume ResumeObject
	err := json.NewDecoder(r.Body).Decode(&resume)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Add to localDB
	id := uuid.NewV1()
	localDB[id.String()] = resume

	// if no errors
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, id.String())
}
