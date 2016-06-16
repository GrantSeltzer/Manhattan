package main

import (
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

	actions := []string{
		"allow",
		"errno",
		"kill",
		"trace",
		"trap",
	}

	for _, action := range actions {
		parseSysCallFlag(action, "clone", &config)

		for _, syscall := range config.Syscalls {
			if syscall.Action != parseAction(action) {
				t.Error("parseSysCallFlag return wrong output", syscall.Action)
			}
		}
	}
}

func TestParseDefaultAction(t *testing.T) {
	config := seccompProfileForTestingPurposes()

	parseDefaultAction("kill", &config)
	if config.DefaultAction != parseAction("kill") {
		t.Error("parseDefaultAction returned wrong output. Expected:",
			parseAction("kill"), "Got: ", config.DefaultAction)
	}
}
