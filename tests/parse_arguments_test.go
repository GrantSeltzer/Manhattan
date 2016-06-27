package main

import (
	"reflect"
	"testing"

	parse "github.com/grantseltzer/Manhattan/ociseccompgen"
	types "github.com/opencontainers/runtime-spec/specs-go"
)

func TestParseArguments(t *testing.T) {

	x := seccompProfileForTestingPurposes()
	syscall := types.Syscall{
		Name:   "clone",
		Action: types.ActErrno,
	}

	err := parse.ParseSyscallFlag("errno", "clone:1:2:3:NE", &x)
	if err != nil {
		t.Error("Parsing arugments returned an error ", err)
	}

	ArgStruct := types.Arg{
		Index:    uint(1),
		Value:    uint64(2),
		ValueTwo: uint64(3),
		Op:       types.OpNotEqual,
	}

	if !reflect.DeepEqual(x.Syscalls[0], ArgStruct) {
		t.Error("Arguments incorrectly parsed.", syscall.Args[0], ArgStruct)
	}
}
