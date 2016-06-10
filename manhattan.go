package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/docker/engine-api/types"
)

const default_seccomp_profile = "/etc/sysconfig/docker-seccomp-profile.json"

func main() {
	configFile, configError := os.Open(default_seccomp_profile)
	fatalErrorCheck(configError,
		"Error opening default configuration. You can specify a custom default with the -location flag")

	defer configFile.Close()

	var SeccompProfile types.Seccomp

	jsonParser := json.NewDecoder(configFile)
	parseError := jsonParser.Decode(&SeccompProfile)
	fatalErrorCheck(parseError, "Error parsing Configuration File")

	kill := flag.String("kill", "", "Respond to system call with KILL")
	trap := flag.String("trap", "", "Respond to system call with TRAP")
	errno := flag.String("errno", "", "Respond to system call with ERRNO")
	trace := flag.String("trace", "", "Respond to system call with TRACE")
	allow := flag.String("allow", "", "Respond to system call with ALLOW")
	//	remove := flag.String("remove", "", "Remove a syscall ")
	defaultAction := flag.String("default", "errno", "Set the default action")
	location := flag.String("location", userHomeDir(),
		"Set the location for the exported seccomp profile.")
	name := flag.String("name", parseTime(), "Set name of output file")
	flag.Parse()

	parseSysCallFlag("kill", *kill, &SeccompProfile)
	parseSysCallFlag("trap", *trap, &SeccompProfile)
	parseSysCallFlag("errno", *errno, &SeccompProfile)
	parseSysCallFlag("trace", *trace, &SeccompProfile)
	parseSysCallFlag("allow", *allow, &SeccompProfile)
	parseDefaultAction(*defaultAction, &SeccompProfile)

	b, marshallError := json.MarshalIndent(SeccompProfile, "", "    ")
	fatalErrorCheck(marshallError, "Error creating Seccomp Profile")

	fullPath := parseLocation(*location, *name)
	fmt.Println(string(b))
	newConfigFile, newConfigError := os.Create(fullPath)
	fatalErrorCheck(newConfigError, "Error creating config file")
	_, writeError := newConfigFile.Write(b)
	fatalErrorCheck(writeError, "Error writing config to file")

}
