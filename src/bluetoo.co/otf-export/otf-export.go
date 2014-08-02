package main

import (
	"flag"
	"fmt"
	"unicode/utf8"
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

func otfExport() bool {
	// parameters
	flag.Parse();
	
	// validation
	switch service {
	case "wunderlist":
	// case "todoist":
	default:
		exitUsage("invalid service")
		return false
	}

	if(utf8.RuneCountInString(username) == 0 || utf8.RuneCountInString(password) == 0) {
		exitUsage("username / password required")
		return false
	}

	return true
}