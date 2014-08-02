package otflib

import (
	"os"
	"io/ioutil"
	"testing"
)

var tempdir string

func init() {
	tempdir,_ = ioutil.TempDir("", "TestAccountExport")
}

func TestAccountExport(t *testing.T) {
	
	// Simple Account
	var filename = tempdir + "/raw.txt"
	t.Logf("Account export path: %s", filename)
	var provider = Provider{ Name:"rawr" }

	lists := []List{}
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
	
	lists = append(lists, *rawr, *boom)
	var account = Account{ Provider: provider, Inbox: inbox, Lists: lists }
	
	result, err := ExportToFile(account, filename)
	if(!result || err != nil) {
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