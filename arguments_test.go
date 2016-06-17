package main

import (
	"testing"

	"github.com/docker/engine-api/types"
)

func TestParseArguments(t *testing.T) {

	syscall := types.Syscall{
		Name:   "clone",
		Action: types.ActErrno,
	}

	delimArgs := []string{
		"clone",
		"1",
		"2",
		"3",
		"NE",
	}

	parseArguments(true, delimArgs, &syscall)

	ArgStruct := types.Arg{
		Index:    uint(1),
		Value:    uint64(2),
		ValueTwo: uint64(3),
		Op:       types.OpNotEqual,
	}

	if *syscall.Args[0] != ArgStruct {
		t.Error("Arguments incorrectly parsed.", *syscall.Args[0], ArgStruct)
	}

}
