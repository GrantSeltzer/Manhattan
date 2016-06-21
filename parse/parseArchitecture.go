package parse

import (
	"fmt"
	"os"
	"strings"

	"github.com/docker/engine-api/types"
)

// ArchFlag takes the raw string passed with the --arch flag, parses it
// and updates the Seccomp config accordingly
func ArchFlag(architectures string, config *types.Seccomp) {
	var architectureArgs []string

	if strings.Contains(architectures, ",") {
		architectureArgs = strings.Split(architectures, ",")
	} else {
		architectureArgs = append(architectureArgs, architectures)
	}

	var arches []types.Arch
	var correctedArch types.Arch
	for _, arg := range architectureArgs {
		correctedArch = parseArch(arg)
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
}

func parseArch(arch string) types.Arch {
	switch arch {
	case "x86":
		return types.ArchX86
	case "amd64":
		return types.ArchX86_64
	case "x32":
		return types.ArchX32
	case "arm":
		return types.ArchARM
	case "arm64":
		return types.ArchAARCH64
	case "mips":
		return types.ArchMIPS
	case "mips64":
		return types.ArchMIPS64
	case "mips64n32":
		return types.ArchMIPS64N32
	case "mipsel":
		return types.ArchMIPSEL
	case "mipsel64":
		return types.ArchMIPSEL64
	case "mipsel64n32":
		return types.ArchMIPSEL64N32
	case "ppc":
		return types.ArchPPC
	case "ppc64":
		return types.ArchPPC64
	case "ppc64le":
		return types.ArchPPC64LE
	case "s390":
		return types.ArchS390
	case "s390x":
		return types.ArchS390X
	default:
		fmt.Println("[*] Unrecognized architecture", arch)
		os.Exit(-6)
		return types.ArchMIPS
	}
}
