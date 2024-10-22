/*
Ben Shirley 2024
This file contains util functions for use in unit tests.
*/

package test

import (
	"database/sql"
	"fmt"
	"os"

	db "server/server/db"
)

// helper method that resets the database to the state specified in the helper file
func resetDatabase() {
	sqlFilepath := "resources/testdb_setup.sql"
	database_url := os.Getenv("DATABASE_URL")
	var err error
	dbInit, err := sql.Open("sqlite3", database_url)
	if err != nil {
		fmt.Println("error opening database")
		panic("something went wrong connecting to database")
	}

	file, err := os.ReadFile(sqlFilepath)
	if err != nil {
		fmt.Println("error reading file")
		panic(err.Error())
	}
	script := string(file)

	_, err = dbInit.Exec(script)
	if err != nil {
		fmt.Println("error running script")
		panic(err.Error())
	}
	fmt.Println("database reset successfully")
	dbInit.Close()

}

func initConnection() {
	os.Setenv("DATABASE_URL", "resources/test_db.db")
	resetDatabase()
	db.InitDatabase()
}
