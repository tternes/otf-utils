package otflib

import (
	"os"
	"io/ioutil"
	"testing"
)

var tempdir string

func init() {
	tempdir,_ = ioutil.TempDir("", "otflibtest")
}

func TestAccountImport(t *testing.T) {
	
	var filename = tempdir + "/importtest.txt"
	t.Logf("Account import path: %s", filename)

	// write a temporary file
	contents := []byte(`
		= Another List =
		[ ] Something
		[X] Completed item
		[ ] ★ Starred item
		= Did It All =
		[X] Yep
		[X] Oh yeah
		[X] Boom
		= Slacker =
		[ ] One
		[ ] Two
		[ ] Three
		[ ] Four
		[ ] Five
		`)
	err := ioutil.WriteFile(filename, contents, 0644)
	if(err != nil) {
		t.Errorf("Failed to write temp file", err)
	}

	// now import the temp file
	var account = Account{}
	err = ImportFromFile(&account, filename)
	
	if(len(account.Lists) != 3) {
		t.Errorf("Invalid number of lists (%d)", len(account.Lists))
	}
	
	taskCount := len(account.Lists[0].Tasks)
	if(taskCount != 3) {
		t.Errorf("Invalid number of tasks in first list (%d)", taskCount)
		t.Error(account)
	}
	
	if(account.Lists[2].Tasks[0].Name != "One") {
		t.Error("Invalid name in slacker list")
	}

	// completed
	if(account.Lists[0].Tasks[1].IsCompleted != true) {
		t.Error("Failed to parse starred task")
	}
	
	// starred	
	if(account.Lists[0].Tasks[2].IsStarred != true) {
		t.Error("Failed to parse starred task")
	}
	
	for _,task := range account.Lists[1].Tasks {
		if(task.IsCompleted != true) {
			t.Error("Incomplete task in fully-completed list")	
		}
	}
}

func TestAccountExport(t *testing.T) {
	
	// Simple Account
	var filename = tempdir + "/exporttest.txt"
	t.Logf("Account export path: %s", filename)
	var provider = Provider{ Name:"rawr" }

	rawr := NewList("rawr")
	rawrA := NewTask("Something in rawr")
	rawrB := NewTask("Another thing")
	AddTask(rawrA, rawr)
	AddTask(rawrB, rawr)

	boom := NewList("boom")
	boomA := NewTask("Goofiness")
	boomB := NewTask("Coolness")
	AddTask(boomA, boom)
	AddTask(boomB, boom)
	
	inbox := NewList("Inbox")
	inboxTask := NewTask("Quickly do something")
	inboxCompletedTask := NewTask("Did this earlier")
	inboxCompletedTask.SetCompleted(true)
	inboxGoldStar := NewTask("This is important")
	inboxGoldStar.SetStarred(true)
	
	AddTask(inboxTask, inbox)
	AddTask(inboxCompletedTask, inbox)
	AddTask(inboxGoldStar, inbox)

	var account = Account{ Provider: provider, Inbox: inbox }
	
	AddList(rawr, &account)
	AddList(boom, &account)
	
	err := ExportToFile(&account, filename)
	if(err != nil) {
		t.Errorf("ExportToFile failed to export to %s", filename)
		t.Error(err)
	}
	
	// file exists?
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("Failed to create output file %s", filename)
	}
	
	// log contents
	exportedFile,_ := ioutil.ReadFile(filename)
	exportedContents := "file contents:\n" + string(exportedFile)
	t.Log(exportedContents)
}