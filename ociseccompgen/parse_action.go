package ociseccompgen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
	types "github.com/opencontainers/runtime-spec/specs-go"
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

	if correctedAction == config.DefaultAction {
		logrus.Info("Action is already set as default")
		return nil
	}

	for _, syscallArg := range syscallArgs {
		delimArgs := strings.Split(syscallArg, ":")
		argSlice, err := parseArguments(delimArgs)
		if err != nil {
			return err
		}

		newSyscall := newSyscallStruct(delimArgs[0], correctedAction, *argSlice)
		descison, err := decideCourseOfAction(&newSyscall, config.Syscalls)
		if err != nil {
			fmt.Println(err)
			return err
		}
		delimDescison := strings.Split(descison, ":")

		if delimDescison[0] == "nothing" {
			logrus.Info("No action taken: ", newSyscall)
		}

		if delimDescison[0] == "append" {
			config.Syscalls = append(config.Syscalls, newSyscall)
		}

		if delimDescison[0] == "overwrite" {
			indexForOverwrite, err := strconv.ParseInt(delimDescison[1], 10, 32)
			if err != nil {
				return err
			}
			config.Syscalls[indexForOverwrite] = newSyscall
		}
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

func newSyscallStruct(name string, action types.Action, args []types.Arg) types.Syscall {
	syscallStruct := types.Syscall{
		Name:   name,
		Action: action,
		Args:   args,
	}
	return syscallStruct
}
