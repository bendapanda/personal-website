/*
*
Ben Shirley 2024
This file contains the handlers for the /api/comments/ endpoint.
*/
package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	database "server/internal/db"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// returns a json list of all comment ids
func GetAllCommentIds(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comments, err := database.GetAllCommentIds()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	encoding, err := json.Marshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	_, err = w.Write(encoding)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// handles distribution of api requests
func CommentsEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetComment(w, r)
	case "POST":
		CreateComment(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}

// returns a single json object containing a comment
func GetComment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idString := params.Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// If this fails we have to assume it did so because the comment was not in the database
	comment, err := database.GetCommentById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	encoding, err := json.Marshal(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	_, err = w.Write(encoding)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

// adds a comment to the database.
func CreateComment(w http.ResponseWriter, r *http.Request) {
	log.Info("Creating comment")
	if r.Method != "POST" {
		log.Error("something other than POST request made to api/comments found its way to CreateComment")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var comment database.Comment
	err := decoder.Decode(&comment)
	if err != nil {
		log.Error("POST body not in comment format")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = database.CreateComment(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Info("comment added to database")

}

// methods beyond this point need some sort of authentication, so will need to wait until
// i can be bothered with jwt.
func EditComment(w http.ResponseWriter, r *http.Request) {

}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

}
