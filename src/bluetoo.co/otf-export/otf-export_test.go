package main

import (
	"testing"
	// "os"
)

func TestNothing (t *testing.T) {
	t.Log("TODO: enable wunderlist and todoist tests")
}

// func TestWunderlistExport(t *testing.T) {
// 	service = "wunderlist"
// 	username = "thaddeus"
// 	password = "lol"
// 	outputPath = "wunderlist_test.txt"
//
// 	if(otfExport()) {
// 		// ... thumbs up ...
// 	} else {
// 		t.Error("failed to export wunderlist")
// 	}
//
// 	// file exists?
// 	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
// 		t.Errorf("Failed to create output file %s", outputPath)
// 	}
//
// 	// TODO: attempt to parse exported file
// }
//
// func TestTodoistExport(t *testing.T) {
// 	t.Skip("Skipping Todoist test - not implemented")
// 	service = "todoist"
// 	username = "thaddeus"
// 	password = "lol"
//
// 	if(otfExport()) {
// 		// ... thumbs up ...
// 	} else {
// 		t.Error("failed to export todoist")
// 	}
// }