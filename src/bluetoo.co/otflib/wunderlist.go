package otflib

import (
	"encoding/json"
	"net/http"
	"fmt"
	"errors"
	"io/ioutil"
	"sort"
)

var wunderlistApiUrl string = "https://api.wunderlist.com"

// --------------------------------------------------

type WunderlistAccount struct {
	account Account
	wusername string
	wpassword string
	wtoken string
}

type WunderlistLoginResponse struct {
	Token string
}

// --------------------------------------------------
// Lists
// --------------------------------------------------
type WunderlistList struct {
	Id string
	Title string
	Position float64
}
type WunderlistLists []WunderlistList
func (items WunderlistLists) Len() int {
	return len(items)
}

func (items WunderlistLists) Less(i, j int) bool {
	return items[i].Position < items[j].Position
}

func (items WunderlistLists) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// --------------------------------------------------
// Tasks
// --------------------------------------------------
type WunderlistTask struct {
	Title string
	Position float64
	List_id string
	Starred bool
	Completed_at string // date
	Deleted_at string // date
}
type WunderlistTasks []WunderlistTask

func (items WunderlistTasks) Len() int {
	return len(items)
}

func (items WunderlistTasks) Less(i, j int) bool {
	return items[i].Position < items[j].Position
}

func (items WunderlistTasks) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// --------------------------------------------------
func NewWunderlistAccount(username string, password string) *WunderlistAccount {
	return &WunderlistAccount{ wusername: username, wpassword: password}
}

// --------------------------------------------------
// interface methods
// --------------------------------------------------
func (w *WunderlistAccount) LoadServiceAccount() bool {

	if(requestWunderlistToken(w) != nil) {
		fmt.Println("Failed to retrieve token")
		return false
	}

	// TODO: goroutine?
	// token available, request lists
	lists, listError := requestWunderlistLists(w)
	if(listError != nil) {
		fmt.Println("Failed to retrieve lists", listError)
		return false
	}
	
	// request tasks
	tasks, taskError := requestWunderlistTasks(w)
	if(taskError != nil) {
		fmt.Println("Failed to retrieve tasks", taskError)
		return false
	}
	
	sort.Sort(lists)
	sort.Sort(tasks)

	// for _,list := range lists {
	// 	// AddList(NewList(list.Title), w.GetAccount())
	// 	fmt.Println(task)
	// }
	
	// Merge lists together
	account := w.GetAccount()
	for _,wlist := range *lists {

		list := NewList(wlist.Title)
		
		// locate items matching current project
		for _,wtask := range *tasks {
			
			if(wtask.List_id == wlist.Id) {
				task := NewTask(wtask.Title)
				task.SetStarred(wtask.Starred)
				task.SetCompleted(len(wtask.Completed_at) > 0)
				AddTask(task, list)
			}

		}

		AddList(list, account)
	}
	
	return true
}

func (w *WunderlistAccount) GetAccount() *Account {
	return &(w.account)
}

func (w *WunderlistAccount) AddListToService(l *List) error {
	return nil
}

func (w *WunderlistAccount) RemoveListFromService(l *List) error {
	return nil
}

// --------------------------------------------------
// Authentication (/login)
// --------------------------------------------------
func requestWunderlistToken(w *WunderlistAccount) error {
	url := wunderlistApiUrl + fmt.Sprintf("/login?email=%s&password=%s", w.wusername, w.wpassword)
	resp, _ := http.Post(url, "text/json", nil)

	switch resp.StatusCode {
	case 200:
	default:
		return errors.New("invalid http response code")
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var loginResponse WunderlistLoginResponse
	json.Unmarshal(body, &loginResponse)
	
	w.wtoken = loginResponse.Token
	return nil
}

// --------------------------------------------------
// Lists
// --------------------------------------------------

func requestWunderlistLists(w *WunderlistAccount) (*WunderlistLists, error) {
	url := wunderlistApiUrl + fmt.Sprintf("/me/lists")
	
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.wtoken))
	req.Header.Set("Content-Type", "text/json")
	resp, _ := client.Do(req)
	
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 200:
	default:
		return nil, errors.New("invalid http response code")
	}
	
	var listsResponse WunderlistLists
	json.Unmarshal(body, &listsResponse)
	
	return &listsResponse, nil
}

// --------------------------------------------------
// Tasks
// --------------------------------------------------
func requestWunderlistTasks(w *WunderlistAccount) (*WunderlistTasks,error) {
	url := wunderlistApiUrl + fmt.Sprintf("/me/tasks")
	
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", w.wtoken))
	req.Header.Set("Content-Type", "text/json")
	resp, _ := client.Do(req)
	
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case 200:
	default:
		return nil, errors.New("invalid http response code")
	}
	
	var tasksResponse WunderlistTasks
	json.Unmarshal(body, &tasksResponse)
	
	return &tasksResponse, nil
}