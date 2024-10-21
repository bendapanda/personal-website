package test

import (
	"errors"
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
	if !comment.Timestamp.Equal(time.Date(2024, 10, 10, 0, 0, 0, 0, time.UTC)) {
		t.Error("The recieved comment has the wrong date.")
	}
}

// Test to ensure error handling works as expected
func TestGetCommentByIdBasicFailure(t *testing.T) {
	_, err := db.GetCommentById(3)
	if err == nil {
		t.Error("There should be an error for a non-existent comment")
	}
	var expectedType *db.NotInDatabaseError
	if !errors.As(err, &expectedType) {
		t.Error("The returned error type is not a NotInDatabaseError")
	}

	if err.Error() != "Object with id 3 not found in Comments" {
		t.Error("The returned error has the wrong message")
	}

}

// Each time a comment is returned, it should be a pointer to the exact same piece of memory
func TestGetCommentByIdCaching(t *testing.T) {
	initConnection()

	comment, err := db.GetCommentById(0)
	if err != nil {
		t.Error("Noting should be wrong here")
	}

	sameComment, err := db.GetCommentById(0)
	if err != nil {
		t.Error("Noting should be wrong here")
	}

	if comment != sameComment {
		t.Error("These two comments should point to the exact same memory address.")
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
