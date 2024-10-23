package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	database "server/internal/db"
	"server/internal/testutils"
)

/**
Ben Shirley, October 2024

This file contains all the nessicary tests for the GetProjects endpoint in the server.
*/

func TestProperUsage1(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// First, generate request objects
	req, err := http.NewRequest("GET", "/api/projects", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProjects)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusOK)
	}

	var resultBody []database.Project
	json.Unmarshal(responseRecorder.Body.Bytes(), &resultBody)

	if len(resultBody) != 2 {
		t.Errorf("response expected to have length 2, got %d", len(resultBody))
	}
	if resultBody[0].Name != "project 1" {
		t.Errorf("first element in response expected to be project 1, was %s", resultBody[0].Name)
	}
	database.CloseConnection()
}

// Test to ensure non-get requests return 400 error codes
func TestProjectsImproperUsage1(t *testing.T) {
	testutils.InitTestConnection()
	database.InitDatabase()

	// First, generate request objects
	req, err := http.NewRequest("PUT", "/api/projects", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetProjects)

	// Now, prompt the server for http results.
	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("response got wrong error code: got %v, expected %v", status, http.StatusBadRequest)
	}
	database.CloseConnection()
}
