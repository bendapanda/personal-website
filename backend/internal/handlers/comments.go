/*
*
Ben Shirley 2024
This file contains the handlers for the /api/comments/ endpoint.
*/
package handlers

import (
	"encoding/json"
	"net/http"
	database "server/internal/db"
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
