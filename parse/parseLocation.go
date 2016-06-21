package parse

import (
	"os"
	"os/user"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

// DefaultFullPath returns the default full path/name for output configuration files
func DefaultFullPath() string {
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
	if err != nil {
		log.Fatal("Could not obtain users home directory. Try setting a custom output location with -location")
	}
	return usr.HomeDir
}

func pwd() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Could not get current working directory")
	}
	return pwd
}
