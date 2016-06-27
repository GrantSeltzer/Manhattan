package ociseccompgen

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	types "github.com/opencontainers/runtime-spec/specs-go"
)

/*******************************************************
This file is a bunch of helper functions for determining
what action should be taken for new syscalls.
See github.com/grantseltzer/manhattan/Resources/Logic.png
for a flowchart and explanation
********************************************************/

func decideCourseOfAction(newSyscall *types.Syscall, syscalls []types.Syscall) (string, error) {
	ruleForSyscallAlreadyExists := false

	var sliceOfDeterminedActions []string
	for i, syscall := range syscalls {
		if syscall.Name == newSyscall.Name {
			ruleForSyscallAlreadyExists = true

			if identical(newSyscall, &syscall) {
				sliceOfDeterminedActions = append(sliceOfDeterminedActions, "nothing")
			}

			if sameAction(newSyscall, &syscall) {
				if bothHaveArgs(newSyscall, &syscall) {
					sliceOfDeterminedActions = append(sliceOfDeterminedActions, "append")
				}
				if onlyOneHasArgs(newSyscall, &syscall) {
					if firstParamOnlyHasArgs(newSyscall, &syscall) {
						sliceOfDeterminedActions = append(sliceOfDeterminedActions, "overwrite:"+strconv.Itoa(i))
					} else {
						sliceOfDeterminedActions = append(sliceOfDeterminedActions, "nothing")
					}
				}
			}

			if !sameAction(newSyscall, &syscall) {
				if bothHaveArgs(newSyscall, &syscall) {
					if sameArgs(newSyscall, &syscall) {
						sliceOfDeterminedActions = append(sliceOfDeterminedActions, "overwrite:"+strconv.Itoa(i))
					}
					if !sameArgs(newSyscall, &syscall) {
						sliceOfDeterminedActions = append(sliceOfDeterminedActions, "append")
					}
				}
				if onlyOneHasArgs(newSyscall, &syscall) {
					sliceOfDeterminedActions = append(sliceOfDeterminedActions, "append")
				}
				if neitherHasArgs(newSyscall, &syscall) {
					sliceOfDeterminedActions = append(sliceOfDeterminedActions, "overwrite:"+strconv.Itoa(i))
				}
			}
		}
	}

	if !ruleForSyscallAlreadyExists {
		sliceOfDeterminedActions = append(sliceOfDeterminedActions, "append")
	}

	// Nothing has highest priority
	for _, determinedAction := range sliceOfDeterminedActions {
		if determinedAction == "nothing" {
			return determinedAction, nil
		}
	}

	// Overwrite has second highest priority
	for _, determinedAction := range sliceOfDeterminedActions {
		if strings.Contains(determinedAction, "overwrite") {
			return determinedAction, nil
		}
	}

	// Append has the lowest priority
	for _, determinedAction := range sliceOfDeterminedActions {
		if determinedAction == "append" {
			return determinedAction, nil
		}
	}

	return "error", fmt.Errorf("Trouble determining action: %s", sliceOfDeterminedActions)
}

func hasArguments(config *types.Syscall) bool {
	nilSyscall := new(types.Syscall)
	if sameArgs(nilSyscall, config) {
		return false
	}
	return true
}

func identical(config1, config2 *types.Syscall) bool {
	if reflect.DeepEqual(config1, config2) {
		return true
	}
	return false
}

func identicalExceptAction(config1, config2 *types.Syscall) bool {
	samename := sameName(config1, config2)
	sameAction := sameAction(config1, config2)
	sameArgs := sameArgs(config1, config2)

	if samename && !sameAction && sameArgs {
		return true
	}
	return false
}

func identicalExceptArgs(config1, config2 *types.Syscall) bool {
	samename := sameName(config1, config2)
	sameAction := sameAction(config1, config2)
	sameArgs := sameArgs(config1, config2)

	if samename && sameAction && !sameArgs {
		return true
	}
	return false
}

func sameName(config1, config2 *types.Syscall) bool {
	if config1.Name == config2.Name {
		return true
	}
	return false
}

func sameAction(config1, config2 *types.Syscall) bool {
	if config1.Action == config2.Action {
		return true
	}
	return false
}

func sameArgs(config1, config2 *types.Syscall) bool {
	if reflect.DeepEqual(config1.Args, config2.Args) {
		return true
	}
	return false
}

func bothHaveArgs(config1, config2 *types.Syscall) bool {
	conf1 := hasArguments(config1)
	conf2 := hasArguments(config2)

	if conf1 && conf2 {
		return true
	}
	return false
}

func onlyOneHasArgs(config1, config2 *types.Syscall) bool {
	conf1 := hasArguments(config1)
	conf2 := hasArguments(config2)

	if (conf1 && !conf2) || (!conf1 && conf2) {
		return true
	}
	return false
}

func neitherHasArgs(config1, config2 *types.Syscall) bool {
	conf1 := hasArguments(config1)
	conf2 := hasArguments(config2)

	if !conf1 && !conf2 {
		return true
	}
	return false
}

func firstParamOnlyHasArgs(config1, config2 *types.Syscall) bool {
	conf1 := hasArguments(config1)
	conf2 := hasArguments(config2)

	if !conf1 && conf2 {
		return true
	}
	return false
}
