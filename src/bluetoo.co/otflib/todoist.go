package otflib

import (
	"encoding/json"
	"net/http"
	"fmt"
	"errors"
	"net/url"
	"io/ioutil"
)

var todoistApiUrl string = "https://todoist.com"

// --------------------------------------------------

type TodoistAccount struct {
	account Account
	tusername string
	tpassword string
	ttoken string
}

type TodoistLoginResponse struct {
	Token string
}

type TodoistSyncResponse struct {
	Projects []TodoistProject
	Items []TodoistItem
}

type TodoistProject struct {
	Name string
	Id int
}

type TodoistItem struct {
	Content string
	Project_id int
	Priority int
}

// --------------------------------------------------

func NewTodoistAccount(username string, password string) *TodoistAccount {
	var provider = Provider{ Name:"todoist" }
	var account = Account{ Provider: provider}
	var todoistAccount = TodoistAccount{ account:account, tusername: username, tpassword: password}
	return &todoistAccount
}

// interface method
func (t *TodoistAccount) LoadServiceAccount() bool {

	if(requestTodoistToken(t) != nil) {
		fmt.Println("Failed to retrieve token")
		return false
	}

	// use the sync API to pull the whole account
	syncResp, syncErr := requestTodoistSyncData(t)
	if(syncErr != nil) {
		fmt.Println("Failed to pull Todoist account data", syncErr)
		return false
	}

	// fill out the OTF Account with projects/tasks
	var account = t.GetAccount()
	for _,project := range syncResp.Projects {

		list := NewList(project.Name)
		
		// locate items matching current project
		for _,item := range syncResp.Items {
			task := NewTask(item.Content)

			// Todoist quirk: priority API values are inverted from their web representation
			switch item.Priority {
			case 2,3,4:
				task.SetStarred(true)
			default:
				task.SetStarred(false)
			}
			
			if(item.Project_id == project.Id) {
				AddTask(task, list)				
			}

		}

		AddList(list, account)
	}

	return true
}

// interface method
func (w *TodoistAccount) GetAccount() *Account {
	return &(w.account)
}

// --------------------------------------------------
// Authentication (/login)
// --------------------------------------------------
func requestTodoistToken(t *TodoistAccount) error {
	loginUrl := todoistApiUrl + fmt.Sprintf("/API/login?email=%s&password=%s", t.tusername, t.tpassword)
	resp, _ := http.Post(loginUrl, "text/json", nil)

	switch resp.StatusCode {
	case 200:
	default:
		return errors.New("invalid http response code")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var loginResponse TodoistLoginResponse
	json.Unmarshal(body, &loginResponse)
	
	t.ttoken = loginResponse.Token
	return nil
}

// --------------------------------------------------
// Sync API
// --------------------------------------------------
func requestTodoistSyncData(t *TodoistAccount) (*TodoistSyncResponse, error) {
	
	syncUrl := todoistApiUrl + fmt.Sprintf("/TodoistSync/v5.3/get")
	resp, _ := http.PostForm(syncUrl, url.Values{ "api_token":{t.ttoken}, "seq_no":{"0"}})

	switch resp.StatusCode {
	case 200:
	default:
		return nil, errors.New("invalid http response code")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var syncResponse TodoistSyncResponse
	json.Unmarshal(body, &syncResponse)	

	return &syncResponse, nil
}