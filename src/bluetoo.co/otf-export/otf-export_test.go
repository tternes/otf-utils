package main

import (
	"testing"
	"os"
)

func TestWunderlistExport(t *testing.T) {

	service = "wunderlist"
	username = os.Getenv("WUNDERLIST_USERNAME")
	password = os.Getenv("WUNDERLIST_PASSWORD")
	outputPath = "wunderlist_test.txt"

	t.Logf("credentials from environment: '%s' / '%s'", username, password)
	
	if(len(username) == 0 || len(password) == 0) {
		t.Skip("Skipping test - required username/password missing")
		return
	}

	if(otfExport() == nil) {
		// ... thumbs up ...
		t.Log("successfully exported account")
	} else {
		t.Error("failed to export wunderlist")
	}

	// file exists?
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Failed to create output file %s", outputPath)
	}

	// TODO: attempt to parse exported file
}

func TestTodoistExport(t *testing.T) {

	service = "todoist"
	username = os.Getenv("TODOIST_USERNAME")
	password = os.Getenv("TODOIST_PASSWORD")
	outputPath = "todoist_test.txt"

	if(otfExport() == nil) {
		// ... thumbs up ...
	} else {
		t.Error("failed to export todoist")
	}
	
	// file exists?
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Failed to create output file %s", outputPath)
	}

	// TODO: attempt to parse exported file
}