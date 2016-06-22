package ociseccompgen

import (
	"fmt"
	"strings"

	types "github.com/opencontainers/runc/libcontainer/configs"
)

// ArchFlag takes the raw string passed with the --arch flag, parses it
// and updates the Seccomp config accordingly
func ArchFlag(architectures string, config *types.Seccomp) error {
	var architectureArgs []string

	if strings.Contains(architectures, ",") {
		architectureArgs = strings.Split(architectures, ",")
	} else {
		architectureArgs = append(architectureArgs, architectures)
	}

	var arches []string
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

func parseArch(arch string) (string, error) {
	switch arch {
	case "x86":
		return "SCMP_ARCH_X86", nil
	case "amd64":
		return "SCMP_ARCH_X86_64", nil
	case "x32":
		return "SCMP_ARCH_X32", nil
	case "arm":
		return "SCMP_ARC_ARM", nil
	case "arm64":
		return "SCMP_ARC_AARCH64", nil
	case "mips":
		return "SCMP_ARC_MIPS", nil
	case "mips64":
		return "SCMP_ARC_MIPS64", nil
	case "mips64n32":
		return "SCMP_ARC_MIPS64N32", nil
	case "mipsel":
		return "SCMP_ARC_MIPSEL", nil
	case "mipsel64":
		return "SCMP_ARC_MIPSEL64", nil
	case "mipsel64n32":
		return "SCMP_ARC_MIPSEL64N32", nil
	case "ppc":
		return "SCMP_ARC_PPC", nil
	case "ppc64":
		return "SCMP_ARC_PPC64", nil
	case "ppc64le":
		return "SCMP_ARC_PPC64LE", nil
	case "s390":
		return "SCMP_ARC_S390", nil
	case "s390x":
		return "SCMP_ARC_S390X", nil
	default:
		return "SCMP_ARC_MIPS", fmt.Errorf("Unrecognized architecutre: %s", arch)
	}
}
