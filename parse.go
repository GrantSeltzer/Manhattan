package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/docker/engine-api/types"
)

func parseSysCallFlag(action string, arguments string, config *types.Seccomp) {

	var syscalls []string
	if strings.Contains(arguments, ",") {
		syscalls = strings.Split(arguments, ",")
	} else if strings.Contains(arguments, "/") {
		syscalls = strings.Split(arguments, "/")
	} else {
		syscalls = append(syscalls, arguments)
	}

	correctedAction := parseAction(action)

	var syscallExists bool = false
	for _, syscall := range syscalls {
		for _, syscallStruct := range config.Syscalls {
			if syscallStruct.Name == syscall {
				syscallExists = true
				syscallStruct.Action = correctedAction
			}
		}
		if syscallExists != true {
			// Add new struct to config.Syscalls
			newSyscallStruct := &types.Syscall{
				Name:   syscall,
				Action: correctedAction}
			config.Syscalls = append(config.Syscalls, newSyscallStruct)
		}
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
