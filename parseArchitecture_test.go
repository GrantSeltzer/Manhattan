package main

import (
	"testing"

	"github.com/docker/engine-api/types"
)

func TestParseArchFlag(t *testing.T) {

	config := seccompProfileForTestingPurposes()

	arches := map[string]types.Arch{
		"x86":         types.ArchX86,
		"amd64":       types.ArchX86_64,
		"x32":         types.ArchX32,
		"arm":         types.ArchARM,
		"arm64":       types.ArchAARCH64,
		"mips":        types.ArchMIPS,
		"mips64":      types.ArchMIPS64,
		"mips64n32":   types.ArchMIPS64N32,
		"mipsel":      types.ArchMIPSEL,
		"mipsel64":    types.ArchMIPSEL64,
		"mipsel64n32": types.ArchMIPSEL64N32,
		"ppc":         types.ArchPPC,
		"ppc64":       types.ArchPPC64,
		"ppc64le":     types.ArchPPC64LE,
		"s390":        types.ArchS390,
		"s390x":       types.ArchS390X,
	}

	for k, v := range arches {
		parseArchFlag(k, &config)
		if config.Architectures[0] != v {
			t.Error("Architectures mismatched", config.Architectures[0], v)
		}
	}

}
