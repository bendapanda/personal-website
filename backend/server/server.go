package main

import (
	"net/http"
	"os"

	handlers "server/server/handlers"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	db "server/server/db"
	utils "server/server/utils"

	"github.com/joho/godotenv"
)

func main() {
	// initialise environment variables
	godotenv.Load()

	// initialise logging
	loggingFile, exists := os.LookupEnv("LOG_FILE")
	if !exists {
		loggingFile = "console"
	}

	utils.InitLogging(loggingFile)

	// set up database
	db.InitDatabase()

	// start router and server
	log.Info("init server")
	PORT := ":8080"
	log.WithField("port", PORT).Info("Starting server on port " + PORT)

	router := mux.NewRouter()

	router.HandleFunc("/api/projects", handlers.GetProjects)

	log.Fatal(http.ListenAndServe(PORT, router))
}
