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

// Tests to ensure GetAllCommentIds works as expected
func TestGetAllComments1(t *testing.T) {
	initConnection()

	comments, err := db.GetAllCommentIds()
	if err != nil {
		t.Error("There should be no error when the function is called correctly. Got", err.Error())
	}
	if len(comments) != 2 {
		t.Error("expected 2 comments, got", len(comments))
	}

	comment0, err := db.GetCommentById(comments[0])
	if err != nil {
		t.Error("either the id in the returned list is incorrect or something went wrong in GetCommentById")
	}

	comment1, err := db.GetCommentById(comments[1])
	if err != nil {
		t.Error("either the id in the returned list is incorrect or something went wrong in GetCommentById")
	}

	if !comment1.Timestamp.Before(comment0.Timestamp) {
		t.Error("comments should be ordered by date")
	}
}
