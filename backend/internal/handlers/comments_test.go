/*
Ben Shirley 2024
This file contains tests for the comment api handlers located at /api/comments/
*/

package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	database "server/internal/db"
	"server/internal/testutils"
	"testing"
)

// Tests proper usage
func TestGetCommentProperUsage1(t *testing.T) {
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
	if resultBody[0] != 1 {
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
