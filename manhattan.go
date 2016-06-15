package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/docker/engine-api/types"
)

const defaultSeccompProfile = "/etc/sysconfig/manhattan.json"

func main() {

	input := flag.String("input", defaultSeccompProfile, "Specify location of base configuration file")
	kill := flag.String("kill", "", "Respond to system call with KILL")
	trap := flag.String("trap", "", "Respond to system call with TRAP")
	errno := flag.String("errno", "", "Respond to system call with ERRNO")
	trace := flag.String("trace", "", "Respond to system call with TRACE")
	allow := flag.String("allow", "", "Respond to system call with ALLOW")
	remove := flag.String("remove", "", "Remove a syscall ")
	defaultAction := flag.String("default", "errno", "Set the default action")
	arch := flag.String("arch", defaultArchitecture(), "Set supported architectures")
	name := flag.String("name", defaultFullPath(), "Set name of output file")

	flag.Parse()

	var SeccompProfile types.Seccomp

	configFile, configError := os.Open(*input)
	if configError != nil {
		fmt.Println("[*] Could not open seccomp profile at", *input)
		fmt.Println("[*] Creating new Profile")
	} else {
		jsonParser := json.NewDecoder(configFile)
		parseError := jsonParser.Decode(&SeccompProfile)
		fatalErrorCheck(parseError, "Error parsing Configuration File")
	}
	defer configFile.Close()

	parseSysCallFlag("kill", *kill, &SeccompProfile)
	parseSysCallFlag("trap", *trap, &SeccompProfile)
	parseSysCallFlag("errno", *errno, &SeccompProfile)
	parseSysCallFlag("trace", *trace, &SeccompProfile)
	parseSysCallFlag("allow", *allow, &SeccompProfile)
	parseDefaultAction(*defaultAction, &SeccompProfile)
	parseArchFlag(*arch, &SeccompProfile)
	removeAction(*remove, &SeccompProfile)

	b, marshallError := json.MarshalIndent(SeccompProfile, "", "    ")
	fatalErrorCheck(marshallError, "Error creating Seccomp Profile")

	newConfigFile, newConfigError := os.Create(*name)
	fatalErrorCheck(newConfigError, "Error creating config file")
	_, writeError := newConfigFile.Write(b)
	fatalErrorCheck(writeError, "Error writing config to file")

}
