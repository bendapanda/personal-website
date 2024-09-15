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
	Started     time.Time
	Finished    sql.NullTime
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
		if err := rows.Scan(&proj.Name, &proj.Description, &proj.URL, &proj.Started, &proj.Finished); err != nil {
			return projects, err
		}
		projects = append(projects, proj)
	}

	if err = rows.Err(); err != nil {
		return projects, err
	}
	return projects, nil
}
