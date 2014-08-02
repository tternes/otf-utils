package otflib

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ExportToFile(account Account, path string) (bool, error) {

	// --------------------------------------------------------------
	// = List =
	// [ ] Incomplete item
	// [X] Completed item
	//
	//
	// = Another List =
	// [ ] Something [2014-01-01]
	// [ ] Do thing
	// [ ] ★ Starred item
	// [ ] Something repeats every seven days [2014-01-01,r7d]
	// --------------------------------------------------------------

	// Procedure:
	// Account
	// foreach list, inbox
	//   foreach task

	var output string;

	// inbox
	if(account.Inbox != nil) {
		appendList(account.Inbox, &output)
	}
	
	for _,list := range account.Lists {
		appendList(&list, &output)
	}
	
	data := []byte(output)
	err := ioutil.WriteFile(path, data, 0644)
	
	if(err != nil) {
		return false, err
	}
	
	return true, nil
}

// Reads the contents of the specified path, returning a boolean for success and the contents of the file as an OTF account
func readOtfFile(path string) (bool, *Account) {
	
	// exists?
	if _, err := os.Stat(path); os.IsNotExist(err) {
	    fmt.Printf("no such file or directory: %s", path)
		return false, nil
	}
	
	// read contents
	fileContents, err := ioutil.ReadFile(path);
	if(err != nil) {
		return false, nil
	}
	
	// parse the contents
	fileContents = fileContents // appease "unused"
	return false, nil
}


func appendList(list *List, output *string) {
	*output += ("= " + list.Name + " =\n")
	for _,task := range list.Tasks {
		appendTask(&task, output)
	}
	*output += "\n"
}

func appendTask(task *Task, output *string) {
	complete := " "
	starred := ""
	if(task.IsCompleted) {
		complete = "X"
	}
	if(task.IsStarred) {
		starred = "★ "
	}
	*output += fmt.Sprintf("[%s] %s%s\n", complete, starred, task.Name)
}