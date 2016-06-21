package main

import (
	"Manhattan/parse"
	"testing"

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
		parse.SysCallFlag(k, "clone", &config)

		for _, syscall := range config.Syscalls {
			if syscall.Action != v {
				t.Error("parseSysCallFlag returned wrong output", syscall.Action, v)
			}
		}
	}
}

func TestParseDefaultAction(t *testing.T) {
	config := seccompProfileForTestingPurposes()

	parse.DefaultAction("kill", &config)
	if config.DefaultAction != types.ActKill {
		t.Error("parseDefaultAction returned wrong output. Expected:",
			types.ActKill, "Got: ", config.DefaultAction)
	}
}
