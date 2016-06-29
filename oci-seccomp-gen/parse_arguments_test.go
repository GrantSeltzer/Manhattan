package ociseccompgen

import (
	"reflect"
	"testing"

	types "github.com/opencontainers/runtime-spec/specs-go"
)

func TestParseArguments(t *testing.T) {
	config := seccompProfileForTestingPurposes()

	err := ParseSyscallFlag("errno", "clone:1:2:3:NE", &config)
	if err != nil {
		t.Error("Parsing arugments returned an error ", err)
	}

	ArgStruct := types.Arg{
		Index:    uint(1),
		Value:    uint64(2),
		ValueTwo: uint64(3),
		Op:       types.OpNotEqual,
	}

	for _, syscall := range config.Syscalls {
		if syscall.Name == "clone" {
			for _, argStruct := range syscall.Args {
				if !reflect.DeepEqual(argStruct, ArgStruct) {
					t.Error("Arguments incorrectly parsed.")
				}
			}
		}
	}
}
