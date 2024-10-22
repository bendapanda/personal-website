/*
Ben Shirley 2024
This file contains tests for the comment api handlers located at /api/comments/
*/

package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/server/handlers"
	"testing"
)

// Tests proper usage
func TestGetCommentProperUsage1(t *testing.T) {
	initConnection()

	// First, generate request objects
	req, err := http.NewRequest("GET", "/api/comments/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetAllCommentIds)

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
}

// Test to ensure non-get requests return 400 error codes
func TestGetAllCommentsImproperUsage1(t *testing.T) {
	initConnection()

	// First, generate request objects
	req, err := http.NewRequest("PUT", "/api/comments/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetAllCommentIds)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}
}
