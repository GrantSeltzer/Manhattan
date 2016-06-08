package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docker/engine-api/types"
)

func parseFlagOpt(action string, arguments string, config *types.Seccomp) {

	var syscalls []string
	if strings.Contains(arguments, ",") {
		syscalls = strings.Split(arguments, ",")
	}
	if strings.Contains(arguments, "/") {
		syscalls = strings.Split(arguments, "/")
	}

	correctedAction := parseAction(action)

	for _, syscall := range syscalls {
		for _, syscallStruct := range config.Syscalls {
			if syscallStruct.Name == syscall {
				syscallStruct.Action = correctedAction
			}
		}
	}

	// Return some type of error?
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
