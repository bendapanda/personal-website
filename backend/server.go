package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/go-co-op/gocron/v2"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"

	db "server/internal/db"
	handlers "server/internal/handlers"
	utils "server/internal/utils"

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

	// Compile cv to html and pdf
	err := updateCV()
	if err != nil {
		log.Error(err.Error())
	}

	// set up the recurring job of compiling my cv to html
	scheculer, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err.Error())
	}
	job, err := scheculer.NewJob(gocron.DailyJob(1, gocron.NewAtTimes(gocron.NewAtTime(0, 0, 0))), gocron.NewTask(updateCV))
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Infof("Created a new cronjob with id %d", job.ID())

	// start router and server
	log.Info("init server")
	PORT := ":8080"
	log.WithField("port", PORT).Info("Starting server on port " + PORT)

	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// the server can also provide the files placed in the public directory
	fileServer := http.FileServer(http.Dir("./static/public"))
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fileServer))

	// setting the server to pass api requests to the correct locations
	router.HandleFunc("/api/projects", handlers.GetProjects)
	router.HandleFunc("/api/comments/all", handlers.GetAllCommentIds)
	router.HandleFunc("/api/comments", handlers.CommentsEndpoint)
	router.HandleFunc("/api/cv", handlers.GetCvHTML)
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(PORT, handler))
}

// compiles the cv located in private to a pdf, and also to html
func updateCV() error {
	cvPath := "./static/private/cv.tex"
	// cv.pdf should be able to be accessed from the static file server, but cv.html should not
	// TODO: fix the scripting to store this file in the correct location
	cvOutputPath := "./static/public/resources"
	mkHTMLCmd := exec.Command("/bin/bash", "./scripts/compile_cv.sh", cvPath, cvOutputPath)
	if err := mkHTMLCmd.Run(); err != nil {
		log.Errorf("Command execution failed: %v", err)
		return err
	}

	log.Info("cv successfully compiled to html and pdf")
	return nil
}
