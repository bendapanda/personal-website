/*
Ben Shirley 2024
This file contains tests for the comment api handlers located at /api/comments/
*/

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	database "server/internal/db"
	"server/internal/testutils"
	"testing"
	"time"
)

// Tests proper usage
func TestGetAllCommentsProperUsage1(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// First, generate request objects
	req, err := http.NewRequest("GET", "/api/comments/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllCommentIds)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusOK)
	}

	var resultBody []int
	json.Unmarshal(responseRecorder.Body.Bytes(), &resultBody)

	if len(resultBody) != 2 {
		t.Errorf("response expected to have length 2, got %d", len(resultBody))
	}
	if resultBody[0] != 0 {
		t.Errorf("first element in response expected to be 1, was %d", resultBody[0])
	}
	database.CloseConnection()
}

// Test to ensure non-get requests return 400 error codes
func TestGetAllCommentsImproperUsage1(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// First, generate request objects
	req, err := http.NewRequest("PUT", "/api/comments/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllCommentIds)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}
	database.CloseConnection()
}

// Test to ensure that we can make a GET request to get a particular comment
func TestGetCommentProperUsage1(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// First, generate request objects
	req, err := http.NewRequest("GET", "/api/comments", nil)
	if err != nil {
		t.Fatal(err)
	}
	query := req.URL.Query()
	query.Add("id", "0")
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetComment)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusOK)
	}
	var result database.Comment
	json.Unmarshal(responseRecorder.Body.Bytes(), &result)

	//ensure we get the correct result
	if result.Id != 0 {
		t.Errorf("response got incorrect id, expected 0, got %d", result.Id)
	}

	database.CloseConnection()

}

// Test to ensure we error if no parameters in request
func TestGetCommentNoParameters(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// First, generate request objects
	req, err := http.NewRequest("GET", "/api/comments", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetComment)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}

	database.CloseConnection()

}

// Test to ensure we error if invalid request parameters.
func TestGetCommentInvalidParameters(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// First, generate request objects
	req, err := http.NewRequest("GET", "/api/comments", nil)
	if err != nil {
		t.Fatal(err)
	}
	query := req.URL.Query()
	query.Add("id", "20")
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetComment)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}

	database.CloseConnection()
}

// Test to ensure that create comment adds it to the database
func TestCreateCommentProperUseage(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()
	// set the comment

	comment := database.Comment{Id: 40, Commenter: "ben", Content: "hello world", Email: "benshirley@hotmail.com", Timestamp: time.Now()}
	encoding, err := json.Marshal(comment)
	if err != nil {
		t.Error("Something went wrong encoding comment in json.")
	}

	// generate request objects
	req, err := http.NewRequest("POST", "/api/comments", bytes.NewReader(encoding))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateComment)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	// We expect a 201 response code
	if status := responseRecorder.Code; status != http.StatusCreated {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusCreated)
	}

	database.CloseConnection()

}

// Test to ensure that create comment errors with no request body
func TestCreateCommentNoBody(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// generate request objects
	req, err := http.NewRequest("POST", "/api/comments", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateComment)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	// We expect a 201 response code
	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}

	database.CloseConnection()

}

// Test to ensure that create comment fails if extra parameters are given
func TestCreateCommentExtraParameters(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()
	// set the comment
	comment := fmt.Sprintf("{Id: 64, Commenter: \"Ben\", Content: \"test\", Email: \"bensemail\", Timestamp: \"%s\", ExtraParam: \"hi ther\"}",
		time.Now().String())

	// generate request objects
	req, err := http.NewRequest("POST", "/api/comments", bytes.NewReader([]byte(comment)))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateComment)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	// We expect a 201 response code
	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}

	database.CloseConnection()

}

// Test to ensure that create comment fails if parameters are missing
func TestCreateCommentMissingParams(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()
	// set the comment
	comment := fmt.Sprintf("{Id: 64, Commenter: \"Ben\", Email: \"bensemail\", Timestamp: \"%s\"}",
		time.Now().String())

	// generate request objects
	req, err := http.NewRequest("POST", "/api/comments", bytes.NewReader([]byte(comment)))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateComment)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	// We expect a 201 response code
	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}

	database.CloseConnection()

}
