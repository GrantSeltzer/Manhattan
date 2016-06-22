package main

import (
	"testing"

	"github.com/grantseltzer/Manhattan/parse"

	"github.com/docker/engine-api/types"
)

func seccompProfileForTestingPurposes() types.Seccomp {
	var syses []*types.Syscall

	sys := types.Syscall{
		Name:   "clone",
		Action: types.ActTrace,
	}
	syses = append(syses, &sys)

	var arches []types.Arch
	arches = append(arches, types.ArchMIPS)

	config := types.Seccomp{
		DefaultAction: types.ActAllow,
		Architectures: arches,
		Syscalls:      syses,
	}

	return config
}

func TestParseSysCallFlag(t *testing.T) {

	config := seccompProfileForTestingPurposes()

	actions := map[string]types.Action{
		"allow": types.ActAllow,
		"errno": types.ActErrno,
		"kill":  types.ActKill,
		"trace": types.ActTrace,
		"trap":  types.ActTrap,
	}

	for k, v := range actions {
		err := parse.SysCallFlag(k, "clone", &config)
		if err != nil {
			t.Error("Parsing syscall flag returned an error ", err)
		}

		for _, syscall := range config.Syscalls {
			if syscall.Action != v {
				t.Error("parseSysCallFlag returned wrong output ", syscall.Action, v)
			}
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
