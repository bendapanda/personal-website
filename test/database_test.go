package test

import (
	"os"
	db "server/server/db"
	"testing"
	"time"
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

// Tests to ensure comments can be retrieved by id
func TestGetCommentByIdBasicSuccess(t *testing.T) {
	initConnection()

	comment, err := db.GetCommentById(0)
	if err != nil {
		t.Error("There should be no error getting a comment: ", err.Error())
	}

	// test to ensure all comment fields are as expected
	if comment.Id != 0 {
		t.Error("The recieved comment should have id 0")
	}
	if comment.Commenter != "test commenter" {
		t.Error("The recieved comment should be by test commenter")
	}
	if comment.Content != "test content" {
		t.Error("The recieved comment should have content test content")
	}
	if comment.Email != "no email" {
		t.Error("The recieved comment should have no email")
	}
	if !comment.Timestamp.Equal(time.Date(2024, 10, 10)) {
		t.Error("The recieved comment has the wrong date.")
	}
}

// test caching for all methods

// Tests to ensure GetAllComments works as expected
func TestGetAllComments1(t *testing.T) {
	initConnection()

	comments, err := db.GetAllComments()
	if err != nil {
		t.Error("There should be no error when the function is called correctly. Got", err.Error())
	}
	if len(comments) != 2 {
		t.Error("expected 2 comments, got", len(comments))
	}

	// The comments should be in order of most recent first.
	comment0 := *comments[0]
	comment1 := *comments[1]
	if !comment0.Timestamp.Before(comment1.Timestamp) {
		t.Error("These comments should ordered with most recent first.")
	}

	if comment0.Commenter != "test commenter 2" {
		t.Error("The first comment is not as expected")
	}

}
