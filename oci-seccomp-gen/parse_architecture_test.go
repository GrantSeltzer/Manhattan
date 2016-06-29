package ociseccompgen

import (
	"reflect"
	"testing"

	types "github.com/opencontainers/runtime-spec/specs-go"
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

	emptyArches := []types.Arch{}

	for k, v := range arches {
		compareArches := append(emptyArches, v)
		err := ParseArchitectureFlag(k, &config)
		if err != nil {
			t.Error("Parsing Arch Flag returned an error")
		}
		if !reflect.DeepEqual(compareArches, config.Architectures) {
			t.Error("Architectures mismatched", compareArches, config.Architectures)
		}
	}
}
