package main

import (
	"os"
	"os/user"
	"strings"
	"time"
)

func defaultFullPath() string {
	return (parseLocation(pwd(), parseTime()))
}

func parseLocation(location, name string) string {
	return strings.TrimSuffix(location, "/") + "/" + name + ".json"
}

func parseTime() string {
	return strings.Replace(time.Now().String(), " ", "", -1)
}

func userHomeDir() string {
	usr, err := user.Current()
	fatalErrorCheck(err, "Could not obtain users home directory. Try setting a custom output location with -location")
	return usr.HomeDir
}

func pwd() string {
	pwd, err := os.Getwd()
	fatalErrorCheck(err, "Couldn't get current working directory")
	return pwd
}
