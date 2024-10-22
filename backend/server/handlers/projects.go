/**
Ben Shirley 2024
This is the handler for /api/projects. Returns a json list containing all the projects stored in the database.
*/

package handlers

import (
	"encoding/json"
	"net/http"
	db "server/server/db"

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
	w.WriteHeader(http.StatusOK)
}
