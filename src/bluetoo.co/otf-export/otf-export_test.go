package main

import (
	"testing"
	"io/ioutil"
	"os"
	"bluetoo.co/otflib"
)

var tempdir string
var usernameVar string
var passwordVar string
func init() {
	tempdir,_ = ioutil.TempDir("", "TestAccountExport")
}

// ------------------------------------------------

func TestWunderlistExport(t *testing.T) {

	usernameVar = "WUNDERLIST_USERNAME"
	passwordVar = "WUNDERLIST_PASSWORD"
	service = "wunderlist"
	outputPath = tempdir + "/wunderlist_test.txt"
	
	commonExportTest(t)
}

func TestTodoistExport(t *testing.T) {

	usernameVar = "TODOIST_USERNAME"
	passwordVar = "TODOIST_PASSWORD"
	service = "todoist"
	outputPath = tempdir + "/todoist_test.txt"
	
	commonExportTest(t)
}

// ------------------------------------------------

func commonExportTest(t *testing.T) {
	
	username = os.Getenv(usernameVar)
	password = os.Getenv(passwordVar)
	
	t.Logf("Output path: %s", outputPath)
	t.Logf("credentials from environment: '%s' / '%s'", username, password)
	
	if(len(username) == 0 || len(password) == 0) {
		t.Skip("Skipping test - required username/password missing")
		return
	}

	if(otfExport() == nil) {
		// ... thumbs up ...
	} else {
		t.Error("failed to export account")
	}
	
	// file exists?
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Errorf("Failed to create output file %s", outputPath)
	}

	// cat output file to log
	outputContents, err := ioutil.ReadFile(outputPath)
	if(err != nil) {
		t.Errorf("Failed to read output file %s", outputPath)
	}
	
	sFileContents := string(outputContents)
	t.Logf("---------------- Contents of %s test ----------------", service)
	t.Log(sFileContents)
	
	resultAccount := otflib.NewAccount()
	err = otflib.ImportFromFile(resultAccount, outputPath)
	if(err!= nil) {
		t.Error("Failed to import result file successful", err)
	}
}