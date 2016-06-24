package ociseccompgen

import (
	"fmt"
	"strings"

	types "github.com/opencontainers/runtime-spec/specs-go"
)

// ParseArchitectureFlag takes the raw string passed with the --arch flag, parses it
// and updates the Seccomp config accordingly
func ParseArchitectureFlag(architectures string, config *types.Seccomp) error {

	var architectureArgs []string

	if strings.Contains(architectures, ",") {
		architectureArgs = strings.Split(architectures, ",")
	} else {
		architectureArgs = append(architectureArgs, architectures)
	}

	var arches []types.Arch
	for _, arg := range architectureArgs {
		correctedArch, err := parseArch(arg)
		if err != nil {
			return err
		}
		shouldAppend := true
		for _, alreadySpecified := range config.Architectures {
			if correctedArch == alreadySpecified {
				shouldAppend = false
			}
		}
		if shouldAppend {
			arches = append(arches, correctedArch)
			config.Architectures = arches
		}
	}
	return nil
}

func parseArch(arch string) (types.Arch, error) {
	switch arch {
	case "x86":
		return types.ArchX86, nil
	case "amd64":
		return types.ArchX86_64, nil
	case "x32":
		return types.ArchX32, nil
	case "arm":
		return types.ArchARM, nil
	case "arm64":
		return types.ArchAARCH64, nil
	case "mips":
		return types.ArchMIPS, nil
	case "mips64":
		return types.ArchMIPS64, nil
	case "mips64n32":
		return types.ArchMIPS64N32, nil
	case "mipsel":
		return types.ArchMIPSEL, nil
	case "mipsel64":
		return types.ArchMIPSEL64, nil
	case "mipsel64n32":
		return types.ArchMIPSEL64N32, nil
	case "ppc":
		return types.ArchPPC, nil
	case "ppc64":
		return types.ArchPPC64, nil
	case "ppc64le":
		return types.ArchPPC64LE, nil
	case "s390":
		return types.ArchS390, nil
	case "s390x":
		return types.ArchS390X, nil
	default:
		return types.ArchMIPS, fmt.Errorf("Unrecognized architecutre: %s", arch)
	}
}
