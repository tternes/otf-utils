package otflib

import (
	"encoding/json"
	"net/http"
	"fmt"
	"errors"
	"net/url"
	"io/ioutil"
	"sort"
	"strconv"
)

var todoistApiUrl string = "https://todoist.com"

// --------------------------------------------------
// Account
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
	Projects TodoistProjects
	Items TodoistItems
}

// --------------------------------------------------
// Project
// --------------------------------------------------

type TodoistProject struct {
	Name string
	Item_order int
	Id int
}
type TodoistProjects []TodoistProject

func (items TodoistProjects) Len() int {
	return len(items)
}

func (items TodoistProjects) Less(i, j int) bool {
	return items[i].Item_order < items[j].Item_order
}

func (items TodoistProjects) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// --------------------------------------------------
// Item
// --------------------------------------------------

type TodoistItem struct {
	Content string
	Project_id int
	Priority int
	Item_order int
}
type TodoistItems []TodoistItem

func (items TodoistItems) Len() int {
	return len(items)
}

func (items TodoistItems) Less(i, j int) bool {
	return items[i].Item_order < items[j].Item_order
}

func (items TodoistItems) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// --------------------------------------------------

func NewTodoistAccount(username string, password string) *TodoistAccount {
	var provider = Provider{ Name:"todoist" }
	var account = Account{ Provider: provider}
	var todoistAccount = TodoistAccount{ account:account, tusername: username, tpassword: password}
	return &todoistAccount
}

// --------------------------------------------------
// interface methods
// --------------------------------------------------
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
	
	// Sort items
	sort.Sort(syncResp.Projects)
	sort.Sort(syncResp.Items)
	
	for _,project := range syncResp.Projects {

		list := NewList(project.Name)
		list.VendorListId = strconv.Itoa(project.Id)
		
		fmt.Println("Loaded project id, name", list.VendorListId, list.Name)
		
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

func (t *TodoistAccount) GetAccount() *Account {
	return &(t.account)
}

func (t *TodoistAccount) AddListToService(l *List) error {
	err := t.requestAddProject(l)
	if(err != nil) {
		return err
	}
	
	for _,task := range l.Tasks {
		
		// don't upload completed tasks
		if(task.IsCompleted == false) {
			err = t.requestAddTask(task, l)
			if(err != nil) {
				return err
			}			
		}
	}

	return nil
}

func (t *TodoistAccount) RemoveListFromService(l *List) error {
	return t.requestRemoveProject(l)
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

// --------------------------------------------------
// Projects
// --------------------------------------------------
func (t *TodoistAccount)requestAddProject(l *List) error {
	// API/addProject?name=..&token=..

	v := url.Values{}
	v.Set("name", l.Name)
	v.Set("token", t.ttoken)

	addUrl  := todoistApiUrl + fmt.Sprintf("/API/addProject?%s", v.Encode())
	resp, _ := http.Get(addUrl)

	switch resp.StatusCode {
	case 200:
	default:
		return errors.New("invalid http response code")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	
	var resultProject TodoistProject
	json.Unmarshal(body, &resultProject)
	
	l.VendorListId = strconv.Itoa(resultProject.Id)

	return nil
}

func (t *TodoistAccount)requestRemoveProject(l *List) error {
	// API/deleteProject?project_id=..&token=..

	v := url.Values{}
	v.Set("project_id", l.VendorListId)
	v.Set("token", t.ttoken)

	rmUrl   := todoistApiUrl + fmt.Sprintf("/API/deleteProject?%s", v.Encode())
	resp, _ := http.Get(rmUrl)

	switch resp.StatusCode {
	case 200:
	default:
		return errors.New("invalid http response code")
	}

	defer resp.Body.Close()

	return nil
}

func (t *TodoistAccount)requestAddTask(task *Task, l *List) error {

	v := url.Values{}
	v.Set("content", task.Name)
	v.Set("project_id", l.VendorListId)
	v.Set("token", t.ttoken)
	
	if(task.IsStarred) {
		v.Set("priority", "4")
	} else {
		v.Set("priority", "1")
	}

	rmUrl   := todoistApiUrl + fmt.Sprintf("/API/addItem?%s", v.Encode())
	resp, _ := http.Get(rmUrl)

	switch resp.StatusCode {
	case 200:
	default:
		return errors.New("invalid http response code")
	}

	defer resp.Body.Close()

	return nil
}