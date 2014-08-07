package main

import (
	"testing"
	"io/ioutil"
	"os"
)

var tempdir string
var usernameVar string
var passwordVar string
func init() {
	tempdir,_ = ioutil.TempDir("", "TestAccountImport")
}

// ------------------------------------------------

func TestWunderlistImport(t *testing.T) {

	// usernameVar = "WUNDERLIST_USERNAME"
	// passwordVar = "WUNDERLIST_PASSWORD"
	// service = "wunderlist"
	// inputPath = tempdir + "/wunderlist_test.txt"
	//
	// commonImportTest(t)
}

func TestTodoistExport(t *testing.T) {

	usernameVar = "TODOIST_USERNAME"
	passwordVar = "TODOIST_PASSWORD"
	service = "todoist"
	inputPath = tempdir + "/todoist_test.txt"

	commonImportTest(t)
}

// ------------------------------------------------

func commonImportTest(t *testing.T) {
	username = os.Getenv(usernameVar)
	password = os.Getenv(passwordVar)
	
	t.Logf("Input path (contents generated if missing): %s", inputPath)
	t.Logf("credentials from environment: '%s' / '%s'", username, password)
	
	if(len(username) == 0 || len(password) == 0) {
		t.Skip("Skipping test - required username/password missing")
		return
	}
	
	// input file exists?
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		// generate a file
		// write a temporary file
		t.Logf("Creating temporary file %s", inputPath)
		contents := []byte(`
			= Have Five =
			[ ] One
			[ ] Two
			[ ] Three
			[ ] Four
			[ ] Five
			= Have Five More =
			[ ] ★ One is important
			[ ] Two
			[ ] Three
			[ ] Four
			[ ] Five
			= Have Twelve =
			[ ] 1
			[ ] 2
			[ ] 3
			[ ] 4
			[ ] 5
			[ ] 6
			[ ] 7
			[ ] 8
			[ ] 9
			[ ] 10
			[ ] 11
			[ ] 12
			`)
		err := ioutil.WriteFile(inputPath, contents, 0644)
		if(err != nil) {
			t.Errorf("Failed to write temp file", err)
		}
	}

	if(otfImport() == nil) {
		// ... thumbs up ...
	} else {
		t.Error("failed to import account")
	}

	// TODO: load account back from service and verify
}