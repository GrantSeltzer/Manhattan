package main

import (
	"testing"

	"github.com/grantseltzer/Manhattan/parse"

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

	err := parse.Arguments(delimArgs, &syscall)
	if err != nil {
		t.Error("Parsing arugments returned an error ", err)
	}

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
