/*
This file handles the conversion of my cv to an html document. This document is then passed off to the caller,
which then renders it as it sees fit.

This handler does not cause the html to compile from the latex. That job is only done once per day to save computing resources.
*/

package handlers

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

// writes the html compiled into static/public/cv.html to the response
// returns a 500 error code if this file is not found
// returns 400 error if non-get request is sent
func GetCvHTML(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		log.Info("non-get request made resume handler")
		return
	}

	w.Header().Set("Content-Type", "text/html")

	resumePath := "./static/public/resources/cv.html"
	b, err := os.ReadFile(resumePath)
	if err != nil {
		log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(b)
	log.Info("api/cv: fetching resume")
}
