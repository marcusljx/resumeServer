package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetResume(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)
	resumeID := values["resumeID"]

	if data, ok := localDB[resumeID]; ok {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	} else {
		http.Error(w, fmt.Sprintf("Resume with ID [%s] not found", resumeID), http.StatusNotFound)
	}
}
