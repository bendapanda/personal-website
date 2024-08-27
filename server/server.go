package main

import (
	"html/template"
	"net/http"
)

func getMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func main() {
	PORT := ":8080"

	http.HandleFunc("/", getMain)

	http.ListenAndServe(PORT, nil)
}
