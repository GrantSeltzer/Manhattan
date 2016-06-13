package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"

	"github.com/docker/engine-api/types"
)

func parseSysCallFlag(action string, arguments string, config *types.Seccomp) {

	if arguments == "" {
		return
	}

	var syscallArgs []string
	if strings.Contains(arguments, ",") {
		syscallArgs = strings.Split(arguments, ",")
	} else if strings.Contains(arguments, "/") {
		syscallArgs = strings.Split(arguments, "/")
	} else {
		syscallArgs = append(syscallArgs, arguments)
	}

	correctedAction := parseAction(action)

	var syscallExists bool = false
	var syscallHasArguments bool = false

	for _, syscallArg := range syscallArgs {
		for _, syscallStruct := range config.Syscalls {
			if syscallStruct.Name == syscallArg {
				syscallExists = true
				if syscallStruct.Args != nil {
					syscallHasArguments = true
				} else {
					syscallStruct.Action = correctedAction
				}
			}
		}
		if syscallExists == false || syscallHasArguments == true {
			var emptyArgs []*types.Arg
			emptyArgs = make([]*types.Arg, 0)
			newSyscallStruct := &types.Syscall{
				Name:   syscallArg,
				Action: correctedAction,
				Args:   emptyArgs}
			config.Syscalls = append(config.Syscalls, newSyscallStruct)
		}
		syscallExists = false
		syscallHasArguments = false
	}
}

// Take passed action, return the SCMP_ACT_<ACTION> version of it
func parseAction(action string) types.Action {
	switch action {
	case "kill":
		return types.ActKill
	case "trap":
		return types.ActTrap
	case "errno":
		return types.ActErrno
	case "trace":
		return types.ActTrace
	case "allow":
		return types.ActAllow
	default:
		fmt.Println("Unrecognized action", action)
		os.Exit(-3)
		return types.ActKill
	}
}

func parseDefaultAction(action string, config *types.Seccomp) {
	config.DefaultAction = parseAction(action)
}

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
