package main

import (
	"os/user"
	"strings"
	"time"
)

func parseLocation(location, name string) string {
	return strings.TrimSuffix(location, "/") + "/" + name + ".json"
}

//returns current time and date as a string without any whitespace
func parseTime() string {
	return strings.Replace(time.Now().String(), " ", "", -1)
}

func userHomeDir() string {
	usr, err := user.Current()
	fatalErrorCheck(err, "Could not obtain users home directory. Try setting a custom output location with -location")
	return usr.HomeDir
}
