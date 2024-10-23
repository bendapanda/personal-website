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

}

// methods beyond this point need some sort of authentication, so will need to wait until
// i can be bothered with jwt.
func EditComment(w http.ResponseWriter, r *http.Request) {

}

func DeleteComment(w http.ResponseWriter, r *http.Request) {

}
