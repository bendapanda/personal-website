package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func InitProjectsTable() {

	database_user := os.Getenv("POSTGRES_USER")
	database_passwrd := os.Getenv("POSTGRES_PASSWORD")

	database_url := "postgres://" + database_user + ":" + database_passwrd + "@localhost:1234?sslmode=disable"
	database, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("connected to database")
	database.Query("SELECT * FROM projects")

	database.Close()
}
