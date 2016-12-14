package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetResume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	values := mux.Vars(r)
	resumeID := values["resumeID"]

	if data, ok := localDB[resumeID]; ok {
		fmt.Fprint(w, data)
	} else {
		http.Error(w, fmt.Sprintf("Resume with ID [%s] not found", resumeID), http.StatusNotFound)
	}
}
