package test

import (
	"os"
	db "server/server/db"
	"testing"
)

func initConnection() {
	os.Setenv("DATABASE_URL", "resources/test_db.db")
	db.InitDatabase()
}

// Tests to ensure GetAllProjects returns the correct number of projects
func TestGetAllProjectsNumberProjects(t *testing.T) {
	initConnection()

	projects, err := db.GetAllProjects()
	if err != nil {
		t.Error("we should not get an error here but got: ", err.Error())
	}
	if len(projects) != 2 {
		t.Error("expected 2, got ", len(projects))
	}

	first_project := projects[0]
	if first_project.Name != "project 1" {
		t.Error("expected project 1, got", first_project.Name)
	}
}
