package database

import (
	"database/sql"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Project struct {
	Name        string
	Description string
	URL         string
	Started     string
	Finished    string
	ImageFile   string
}

// Initialises the database connection
func InitDatabase() error {
	database_url := os.Getenv("DATABASE_URL")
	log.Info(database_url)
	var err error
	db, err = sql.Open("sqlite3", database_url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Info("Connected to database")
	return nil
}

// fetches all projects in the database
func GetAllProjects() ([]Project, error) {
	queryString := "SELECT * FROM projects"
	rows, err := db.Query(queryString)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var proj Project
		startTime := time.Now()
		var endTime sql.NullTime

		if err := rows.Scan(&proj.Name, &proj.Description, &proj.URL, &startTime, &endTime, &proj.ImageFile); err != nil {
			return projects, err
		}
		proj.Started = startTime.Format(time.DateOnly)
		if endTime.Valid {
			proj.Finished = endTime.Time.Format(time.DateOnly)
		} else {
			proj.Finished = "present"
		}

		projects = append(projects, proj)
	}

	if err = rows.Err(); err != nil {
		return projects, err
	}
	return projects, nil
}
