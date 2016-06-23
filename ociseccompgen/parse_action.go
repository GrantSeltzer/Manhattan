package ociseccompgen

import (
	"fmt"
	"strings"

	types "github.com/opencontainers/runc/libcontainer/configs"
)

// ParseSyscallFlag takes the name of the action, the arguments (syscalls) that were
// passed with it at the command line and a pointer to the config struct. It parses
// the action and syscalls and updates the config accordingly
func ParseSyscallFlag(action string, arguments string, config *types.Seccomp) error {

	if arguments == "" {
		return nil
	}

	syscallArgs := strings.Split(arguments, ",")

	correctedAction, err := parseAction(action)
	if err != nil {
		return err
	}

	for _, syscallArg := range syscallArgs {
		delimArgs := strings.Split(syscallArg, ":")
		argSlice, err := parseArguments(delimArgs)
		if err != nil {
			return err
		}

		newSyscall := newSyscallStruct(delimArgs[0], correctedAction, argSlice)

		var sysCallAlreadySpecified bool
		for _, syscall := range config.Syscalls {
			sysCallAlreadySpecified = compareSyscalls(&newSyscall, syscall)
		}

		if !sysCallAlreadySpecified {
			config.Syscalls = append(config.Syscalls, &newSyscall)
		}
	}
	return nil
}

// Take passed action, return the SCMP_ACT_<ACTION> version of it
func parseAction(action string) (types.Action, error) {
	switch action {
	case "kill":
		return types.Kill, nil
	case "trap":
		return types.Trap, nil
	case "errno":
		return types.Errno, nil
	case "trace":
		return types.Trace, nil
	case "allow":
		return types.Allow, nil
	default:
		return types.Kill, fmt.Errorf("Unrecognized action: %s", action)

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

func newSyscallStruct(name string, action types.Action, args []*types.Arg) types.Syscall {
	syscallStruct := types.Syscall{
		Name:   name,
		Action: action,
		Args:   args,
	}
	return syscallStruct
}
