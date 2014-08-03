package main

import (
	"errors"
	"flag"
	"fmt"
	"unicode/utf8"
	"bluetoo.co/otflib"
)

var service string
var username string
var password string
var outputPath string

func init() {
	flag.StringVar(&service, "service", "", "Service name (e.g. \"wunderlist\" or \"todoist\")")
	flag.StringVar(&username, "username", "", "Username / email for specified service")
	flag.StringVar(&password, "password", "", "Password for specified service")
	flag.StringVar(&outputPath, "output", service+".txt", "Path for exported file")
}

func exitUsage(message string) {
	fmt.Println("Usage", message)
	// flag.PrintDefaults()
	// os.Exit(-1);
}

func main() {
	otfExport()
}

func otfExport() error {
	// parameters
	flag.Parse()
	
	var account otflib.Service;
	
	// validation
	switch service {
	case "wunderlist":
		account = otflib.NewWunderlistAccount(username, password)
	case "todoist":
		account = otflib.NewTodoistAccount(username, password)
	default:
		exitUsage("invalid service")
		return errors.New("Invalid service")
	}

	if(utf8.RuneCountInString(username) == 0 || utf8.RuneCountInString(password) == 0) {
		exitUsage("username / password required")
		return errors.New("Username and password required")
	}

	// load account
	if(!account.LoadServiceAccount()) {
		return errors.New("Failed to load service account")
	}
	
	// export
	err := otflib.ExportToFile(account.GetAccount(), outputPath)
	return err
}