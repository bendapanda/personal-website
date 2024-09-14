package database

import (
	"database/sql"
	"os"

	log "github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

var db *sql.DB

type project struct {
	name        string
	description string
}

// Initialises the database connection
func InitDatabase() error {
	postgres_username := os.Getenv("POSTGRES_USER")
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	postgres_url := "postgresql://" + postgres_username + ":" + postgres_password + "@localhost:5432?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", postgres_url)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// fetches all projects in the database
func GetAllProjects() ([]project, error) {

	queryString := "SELECT * FROM projects"
	rows, err := db.Query(queryString)
	if err != nil {

	}

}
