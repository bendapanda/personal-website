package database

/*
This file contains tests for the database package. To run the tests, first ensure that the
database is reset correctly by running testdb_setup.sql.

Author: Ben Shirley
2024
*/

import (
	"errors"
	"fmt"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"server/internal/testutils"
)

// Tests to ensure GetAllProjects returns the correct number of projects
func TestGetAllProjectsNumberProjects(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	projects, err := GetAllProjects()
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

	CloseConnection()
}

// Tests to ensure comments can be retrieved by id
func TestGetCommentByIdBasicSuccess(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	comment, err := GetCommentById(1)
	if err != nil {
		t.Error("There should be no error getting a comment: ", err.Error())
	}
	if comment == nil {
		t.Error("something should be returned.")
	}

	// test to ensure all comment fields are as expected
	if comment.Id != 1 {
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
	if !comment.Timestamp.Equal(time.Date(2024, 10, 20, 0, 0, 0, 0, time.UTC)) {
		t.Error("The recieved comment has the wrong date.")
	}
	CloseConnection()
}

// Test to ensure error handling works as expected
func TestGetCommentByIdBasicFailure(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	_, err := GetCommentById(3)
	if err == nil {
		t.Error("There should be an error for a non-existent comment")
	}
	var expectedType *DatabaseError
	if !errors.As(err, &expectedType) {
		t.Error("The returned error type is not a DatabaseError")
	}

	CloseConnection()
}

// Each time a comment is returned, it should be a pointer to the exact same piece of memory
func TestGetCommentByIdCaching(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	comment, err := GetCommentById(0)
	if err != nil {
		t.Error("Noting should be wrong here")
	}

	sameComment, err := GetCommentById(0)
	if err != nil {
		t.Error("Noting should be wrong here")
	}

	if comment != sameComment {
		t.Error("These two comments should point to the exact same memory address.")
	}
	CloseConnection()
}

// Tests to ensure GetAllCommentIds works as expected
func TestGetAllComments1(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	comments, err := GetAllCommentIds()
	if err != nil {
		t.Error("There should be no error when the function is called correctly. Got", err.Error())
	}
	if len(comments) != 2 {
		t.Error("expected 2 comments, got", len(comments))
	}

	comment0, err := GetCommentById(comments[0])
	if err != nil {
		t.Error("either the id in the returned list is incorrect or something went wrong in GetCommentById")
	}

	comment1, err := GetCommentById(comments[1])
	if err != nil {
		t.Error("either the id in the returned list is incorrect or something went wrong in GetCommentById")
	}

	if comment1.Timestamp.Before(comment0.Timestamp) {
		t.Error("comments should be ordered by date")
	}
	CloseConnection()
}

// Test to ensure comments get added to the database as expected.
func TestCreateCommentBasic(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	commentToAdd := Comment{Id: 7, Commenter: "Ben", Email: "test email",
		Content: "This is a new Comment", Timestamp: time.Now()}
	newId, err := CreateComment(&commentToAdd)
	if err != nil {
		t.Error("This is a valid usage of the CreateComment method")
	}

	retrievedComment, err := GetCommentById(newId)
	if err != nil {
		t.Error("Nothing should go wrong here")
	}
	fmt.Println(retrievedComment.Commenter)
	if retrievedComment.Commenter != commentToAdd.Commenter {
		t.Errorf("The recieved comment and the comment we added should have the same fields. got %s, expected %s", retrievedComment.Commenter, commentToAdd.Commenter)
	}

	commentIds, err := GetAllCommentIds()
	if err != nil {
		t.Error("Nothing should go wrong here")
	}
	contains := false
	for _, val := range commentIds {
		if val == newId {
			contains = true
		}
	}
	if !contains {
		t.Error("the list of all comment Ids should contain the id of the comment added")
	}

	CloseConnection()
}

// Test to ensure that we can edit a comment
func TestEditComment(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	existantComment := Comment{Id: 1, Commenter: "ben", Content: "this comment already exists", Timestamp: time.Now()}
	err := EditComment(&existantComment)
	if err != nil {
		t.Error("There should be no problem editing this comment")
	}

	retrievedComment, err := GetCommentById(1)
	if err != nil {
		t.Error("There should be no problem calling this method.")
	}

	if retrievedComment.Commenter != existantComment.Commenter {
		t.Error("incorrect commenter")
	}
	if retrievedComment.Content != existantComment.Content {
		t.Error("incorrect content")
	}
	// importantly, the edit should change the cached comment so that all comments are the same.
	CloseConnection()

}

// Test to ensure that we cannot edit a comment that is not in the database
func TestEditCommentNonExistent(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	existantComment := Comment{Id: 5, Commenter: "ben", Content: "this comment already exists", Timestamp: time.Now()}
	err := EditComment(&existantComment)
	if err == nil {
		t.Error("The comment does not exist so we should get and error")
	}

	var expectedType *DatabaseError
	if !errors.As(err, &expectedType) {
		t.Error("The returned error type is not a DatabaseError")
	}

	if err.Error() != "The comment attempted to edit does not exist in database" {
		t.Error("The returned error has the wrong message")
	}
	CloseConnection()
}

// Test to ensure delete method works as expected.
func TestDeleteComment(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	_, err := GetCommentById(1)
	if err != nil {
		t.Error("There should be no problem getting a comment")
	}

	err = DeleteComment(1)
	if err != nil {
		t.Error("There should be no error deleting a comment that exists")
	}

	results, err := GetAllCommentIds()
	if err != nil {
		t.Error("This is a standard call, should be no error")
	}
	if len(results) != 1 {
		t.Error("There should only be one comment")
	}
	if results[0] != 2 {
		t.Error("The only comment in the database should have id 2")
	}

	_, err = GetCommentById(0)
	if err == nil {
		t.Error("There is no such comment in the database so we should not be able to retrieve it.")
	}

	CloseConnection()
}

// Test DeleteComment method to ensure an error is raised if a non-existant comment is attempted to be deleted
func TestDeleteCommentThatDoesNotExists(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	err := DeleteComment(420)
	if err == nil {
		t.Error("It should be impossible to delete a comment that doesn't exist")
	}

	var expectedType *DatabaseError
	if !errors.As(err, &expectedType) {
		t.Error("The returned error type is not a DatabaseError")
	}
	CloseConnection()
}

// Test to ensure comment rate limits are in place.
func TestRateLimit(t *testing.T) {
	testutils.InitTestConnection()
	InitDatabase()

	var err error
	for i := 0; i < 102; i++ {
		toAdd := Comment{Id: i + 20, Commenter: "Ben", Email: "No Email",
			Content: "New Comment", Timestamp: time.Now()}
		if err == nil {
			_, err = CreateComment(&toAdd)
		}
	}

	if err == nil {
		t.Error("At some point we should see and error, we need rate limits!")
	}
	CloseConnection()
}
