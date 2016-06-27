package main

import (
	"testing"

	parse "github.com/grantseltzer/Manhattan/ociseccompgen"

	types "github.com/opencontainers/runtime-spec/specs-go"
)

func seccompProfileForTestingPurposes() types.Seccomp {
	syses := new([]types.Syscall)

	sys := types.Syscall{
		Name:   "clone",
		Action: types.ActTrace,
	}
	*syses = append(*syses, sys)

	var arches []types.Arch
	arches = append(arches, types.ArchMIPS)

	config := types.Seccomp{
		DefaultAction: types.ActAllow,
		Architectures: arches,
		Syscalls:      *syses,
	}

	return config
}

func TestParseSysCallFlagOne(t *testing.T) {
	config := seccompProfileForTestingPurposes()

	actions := map[string]types.Action{
		"allow": types.ActAllow,
		"errno": types.ActErrno,
		"kill":  types.ActKill,
		"trace": types.ActTrace,
		"trap":  types.ActTrap,
	}

	for k, v := range actions {
		err := parse.ParseSyscallFlag(k, "clone", &config)
		if err != nil {
			t.Error("Parsing syscall flag returned an error ", err)
		}
		if config.Syscalls[0].Action != v && config.DefaultAction != v {
			t.Error("parseSysCallFlag returned wrong output ", config.Syscalls[0].Action, v)
		}
	}
}

func TestParseDefaultAction(t *testing.T) {
	config := seccompProfileForTestingPurposes()

	err := parse.DefaultAction("kill", &config)
	if err != nil {
		t.Error("Parsing default action returned an error ", err)
	}
	if config.DefaultAction != types.ActKill {
		t.Error("parseDefaultAction returned wrong output. Expected:",
			types.ActKill, "Got: ", config.DefaultAction)
	}
}
