package handlers

import (
	"encoding/json"
	"net/http"
	db "server/internal/db"

	log "github.com/sirupsen/logrus"
)

// Returns a json encoding of all projects listed in the database.
func GetProjects(w http.ResponseWriter, r *http.Request) {
	log.Info("Fetching projects")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	projects, err := db.GetAllProjects()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	jsonEncoding, err := json.Marshal(projects)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonEncoding)
}
