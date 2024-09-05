package main

import (
	"flag"
	"net/http"

	handlers "server/server/handlers"

	log "github.com/sirupsen/logrus"

	utils "server/server/utils"
)

func main() {
	loggingFile := flag.String("loggingfile", "console", "the file for logs to be stored in")
	flag.Parse()

	utils.InitLogging(*loggingFile)
	log.Info("init server")
	PORT := ":8080"
	log.WithField("port", PORT).Info("Starting server on port " + PORT)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.GetLanding)

	log.Fatal(http.ListenAndServe(PORT, nil))
}
