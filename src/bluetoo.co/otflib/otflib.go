package otflib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func ExportToFile(account *Account, path string) error {

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
		appendList(list, &output)
	}
	
	data := []byte(output)
	err := ioutil.WriteFile(path, data, 0644)
	
	if(err != nil) {
		return err
	}
	
	return nil
}

const (
	sInitial = 0
	sListName     = 1 // collecting list name
	sTaskStatus   = 2 // waiting or 'X' or ']'
	sTaskStar     = 3 // waiting for ★ or non-whitespace
	sTaskContents = 4 // collecting task contents
)
func ImportFromFile(account *Account, path string) error {

	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	var list *List = nil
	for scanner.Scan() {
		line := scanner.Text()

		collectedString := ""
		state := sInitial
		var task *Task = nil

		for _,_char := range strings.Split(line, "") {
			
			char := string(_char)
			switch state {
			case sInitial:
				switch char {
				case "=":
					state = sListName
				case "[":
					state = sTaskStatus
					task = NewTask("")
				}
				
			case sListName:
				switch char {
				case "=":
					list = NewList(strings.TrimSpace(collectedString))
					AddList(list, account)

					state = sInitial
					collectedString = ""
				default:
					collectedString += char
				}

			case sTaskStatus:
				switch char {
				case "X":
					task.SetCompleted(true)
				case "]":
					state = sTaskStar
				}
				
			case sTaskStar:
				switch char {
				case " ":
				case "★":
					task.SetStarred(true)
				default:
					task.Name += char
					state = sTaskContents
				}
				
			case sTaskContents:
				task.Name += char
			}
		}
		
		if(task != nil && len(task.Name) > 0) {
			
			if(list == nil) {
				panic("invalid list for pending task")
			}

			fmt.Println("Adding task", task)
			AddTask(task, list)
		}
	}

	return nil
}

func appendList(list *List, output *string) {
	*output += ("= " + list.Name + " =\n")
	for _,task := range list.Tasks {
		appendTask(task, output)
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