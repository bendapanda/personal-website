package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	utils "server/server/utils"
)

func getMain(w http.ResponseWriter, r *http.Request) {

	log.Info("fetching main")

	t, err := template.ParseFiles("templates/index.html", "templates/navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	age := time.Now().Year() - 2004
	if time.Now().Month() == time.January && time.Now().Day() < 5 {
		age -= 1
	}

	MainInfo := struct {
		Age string
	}{
		fmt.Sprintf("%d", age),
	}
	t.Execute(w, MainInfo)
}

func main() {
	loggingFile := flag.String("loggingfile", "console", "the file for logs to be stored in")
	flag.Parse()

	utils.InitLogging(*loggingFile)
	log.Info("init server")

	PORT := ":8080"
	fs := http.FileServer(http.Dir("./style"))
	http.Handle("/style/", http.StripPrefix("/style/", fs))

	http.HandleFunc("/", getMain)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
