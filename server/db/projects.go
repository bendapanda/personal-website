package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func InitProjectsTable() {

	database, err := sql.Open("postgres", "user=bensh dbname=personal_website")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("connected to database")

	database.Close()
}
