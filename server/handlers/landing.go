package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	db "server/server/db"

	log "github.com/sirupsen/logrus"
)

// Called when a user requests the landing page. Dynamically assembles the page using results from the database
func GetLanding(w http.ResponseWriter, r *http.Request) {
	// Parse all templates used in page construction
	temp := template.Must(template.ParseFiles("templates/index.html", "templates/navbar.html", "templates/project.html"))

	log.Info("fetching main")

	// parse page data
	age := time.Now().Year() - 2004
	if time.Now().Month() == time.January && time.Now().Day() < 5 {
		age -= 1
	}

	projects, err := db.GetAllProjects()
	if err != nil {
		log.Error(err.Error())
		projects[0] = db.Project{
			Name:        "Oops",
			Description: "Sorry, something went wrong here.\n Looks like Ben needs to fix something...",
			Started:     "never",
			Finished:    "never",
			URL:         "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		}
	}

	// Assemble page info
	MainInfo := struct {
		Age      string
		Projects []db.Project
	}{
		fmt.Sprintf("%d", age),
		projects,
	}

	// Build the page
	if err := temp.Execute(w, MainInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Something went wrong!")
		return
	}
}
