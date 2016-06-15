package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/docker/engine-api/types"
)

func parseSysCallFlag(action string, arguments string, config *types.Seccomp) {

	if arguments == "" {
		return
	}

	var (
		argsSpecified  bool
		syscallArgName string
		delimArgs      []string
	)

	/** Split up syscall specifications **/
	var syscallArgs []string
	if strings.Contains(arguments, ",") {
		syscallArgs = strings.Split(arguments, ",")
	} else {
		syscallArgs = append(syscallArgs, arguments)
	}

	correctedAction := parseAction(action)

	syscallExists := false
	syscallHasArguments := false

	/** For each syscall specified for a specific action**/
	for _, syscallArg := range syscallArgs {

		if strings.Contains(syscallArg, ":") {
			argsSpecified = true
			delimArgs = strings.Split(syscallArg, ":")
			syscallArgName = delimArgs[0]
		} else {
			syscallArgName = syscallArg
		}

		/** Go through the syscalls in the existing config **/
		for _, syscallStruct := range config.Syscalls {
			if syscallStruct.Name == syscallArgName {
				syscallExists = true
				if syscallStruct.Args != nil {
					syscallHasArguments = true
				} else {
					syscallStruct.Action = correctedAction
					parseArguments(argsSpecified, delimArgs, syscallStruct)
				}
			}
		}
		if syscallExists == false || syscallHasArguments == true {
			var emptyArgs []*types.Arg
			emptyArgs = make([]*types.Arg, 0)
			newSyscallStruct := &types.Syscall{
				Name:   syscallArgName,
				Action: correctedAction,
				Args:   emptyArgs}
			parseArguments(argsSpecified, delimArgs, newSyscallStruct)
			config.Syscalls = append(config.Syscalls, newSyscallStruct)

		}
		syscallExists = false
		syscallHasArguments = false
		argsSpecified = false
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
