package ociseccompgen

import (
	"reflect"

	types "github.com/opencontainers/runtime-spec/specs-go"
)

/*******************************************************
This file is a bunch of helper functions for determining
what action should be taken for new syscalls. See flowchart
for more details.
********************************************************/

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

	if conf1 && !conf2 {
		return true
	}
	return false
}
