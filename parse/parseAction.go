package parse

import (
	"fmt"
	"strings"

	"github.com/docker/engine-api/types"
)

//SysCallFlag takes the name of the action, the arguments (syscalls) that were
//passed with it at the command line and a pointer to the config struct. It parses
//the action and syscalls and updates the config accordingly
func SysCallFlag(action string, arguments string, config *types.Seccomp) error {
	if arguments == "" {
		return nil
	}

	var (
		argsSpecified       bool
		syscallArgName      string
		delimArgs           []string
		syscallExists       bool
		syscallHasArguments bool
	)

	/** Split up syscall specifications **/
	var syscallArgs []string
	if strings.Contains(arguments, ",") {
		syscallArgs = strings.Split(arguments, ",")
	} else {
		syscallArgs = append(syscallArgs, arguments)
	}

	correctedAction, err := parseAction(action)
	if err != nil {
		return err
	}

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
					if argsSpecified {
						Arguments(delimArgs, syscallStruct)
					}
				}
			}
		}
		if !syscallExists || syscallHasArguments {
			var emptyArgs []*types.Arg
			emptyArgs = make([]*types.Arg, 0)
			newSyscallStruct := &types.Syscall{
				Name:   syscallArgName,
				Action: correctedAction,
				Args:   emptyArgs}
			if argsSpecified {
				Arguments(delimArgs, newSyscallStruct)
			}
			config.Syscalls = append(config.Syscalls, newSyscallStruct)

		}
		syscallExists = false
		syscallHasArguments = false
		argsSpecified = false
	}
	return nil
}

// Take passed action, return the SCMP_ACT_<ACTION> version of it
func parseAction(action string) (types.Action, error) {
	switch action {
	case "kill":
		return types.ActKill, nil
	case "trap":
		return types.ActTrap, nil
	case "errno":
		return types.ActErrno, nil
	case "trace":
		return types.ActTrace, nil
	case "allow":
		return types.ActAllow, nil
	default:
		return types.ActKill, fmt.Errorf("Unrecognized action: %s", action)

	}
}

//DefaultAction simply sets the default action of the seccomp configuration
func DefaultAction(action string, config *types.Seccomp) error {
	defaultAction, err := parseAction(action)
	if err != nil {
		return err
	}
	config.DefaultAction = defaultAction
	return nil
}
