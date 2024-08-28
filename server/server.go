package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func getMain(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/navbar.html")

	age := time.Now().Year() - 2004
	if time.Now().Month() == time.January && time.Now().Day() < 5 {
		age -= 1
	}

	MainInfo := struct {
		Age string
	}{
		fmt.Sprintf("%d", age),
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, MainInfo)
}

func main() {
	PORT := ":8080"
	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", getMain)

	http.ListenAndServe(PORT, nil)
}
