package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

var temp = template.Must(template.ParseFiles("templates/index.html", "templates/navbar.html", "templates/project.html"))

func GetLanding(w http.ResponseWriter, r *http.Request) {

	log.Info("fetching main")

	age := time.Now().Year() - 2004
	if time.Now().Month() == time.January && time.Now().Day() < 5 {
		age -= 1
	}

	MainInfo := struct {
		Age string
	}{
		fmt.Sprintf("%d", age),
	}
	if err := temp.Execute(w, MainInfo); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Something went wrong!")
		return
	}
}
