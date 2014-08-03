package otflib

import (
	"encoding/json"
	"net/http"
	"fmt"
	"errors"
	"io/ioutil"
)

var wunderlistApiUrl string = "https://api.wunderlist.com"

// --------------------------------------------------

type WunderlistAccount struct {
	Account Account
	wusername string
	wpassword string
	wtoken string
}

type WunderlistLoginResponse struct {
	Token string
}

func NewWunderlistAccount(username string, password string) *WunderlistAccount {
	return &WunderlistAccount{ wusername: username, wpassword: password}
}

// interface method
func (w WunderlistAccount) LoadServiceAccount() bool {

	if(requestWunderlistToken(&w) != nil) {
		fmt.Println("Failed to retrieve token")
		return false
	}

	// token available, request lists
	if(requestWunderlistLists(&w) != nil) {
		fmt.Println("Failed to retrieve lists")
		return false
	}
	
	// foreach list, request tasks
	
	return true
}

// interface method
func (w *WunderlistAccount) GetAccount() *Account {
	return &(w.Account)
}

// --------------------------------------------------
// Authentication (/login)
// --------------------------------------------------
func requestWunderlistToken(w *WunderlistAccount) error {
	url := wunderlistApiUrl + fmt.Sprintf("/login?email=%s&password=%s", w.wusername, w.wpassword)
	resp, _ := http.Post(url, "text/json", nil)
	
	fmt.Println(resp.Status)

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

func requestWunderlistLists(w *WunderlistAccount) error {
	return nil	
}

// --------------------------------------------------
// Tasks
// --------------------------------------------------