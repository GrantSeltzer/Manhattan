package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/docker/engine-api/types"
	"github.com/urfave/cli"
)

const defaultSeccompProfile = "/etc/sysconfig/manhattan.json"

func main() {

	var (
		input         string
		kill          string
		trap          string
		errno         string
		trace         string
		allow         string
		remove        string
		defaultAction string
		arch          string
		name          string
	)

	app := cli.NewApp()
	app.Name = "manhattan"
	app.Usage = "Create Docker compliant seccomp json configurations"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "input, i",
			Value:       defaultSeccompProfile,
			Usage:       "Specify location of base configuration file",
			Destination: &input,
		},
		cli.StringFlag{
			Name:        "kill, k",
			Value:       "",
			Usage:       "Respond to system call with KILL",
			Destination: &kill,
		},
		cli.StringFlag{
			Name:        "trap, p",
			Value:       "",
			Usage:       "Respond to system call with TRAP",
			Destination: &trap,
		},
		cli.StringFlag{
			Name:        "errno, e",
			Value:       "",
			Usage:       "Respond to system call with ERRNO",
			Destination: &errno,
		},
		cli.StringFlag{
			Name:        "trace, c",
			Value:       "",
			Usage:       "Respond to system call with TRACE",
			Destination: &trace,
		},
		cli.StringFlag{
			Name:        "allow, a",
			Value:       "",
			Usage:       "Respond to system call with ALLOW",
			Destination: &allow,
		},
		cli.StringFlag{
			Name:        "remove, r",
			Value:       "",
			Usage:       "Remove a syscall",
			Destination: &remove,
		},
		cli.StringFlag{
			Name:        "default, d",
			Value:       "errno",
			Usage:       "Set the default action for syscalls not specified",
			Destination: &defaultAction,
		},
		cli.StringFlag{
			Name:        "arch, l",
			Value:       "amd64,x86,x32",
			Usage:       "Set supported architectures",
			Destination: &arch,
		},
		cli.StringFlag{
			Name:        "name, n",
			Value:       defaultFullPath(),
			Usage:       "Set name of output file",
			Destination: &name,
		},
	}
	app.Version = version
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		return nil
	}

	app.Run(os.Args)

	var SeccompProfile types.Seccomp

	configFile, configError := os.Open(input)
	if configError != nil {
		fmt.Println("[*] Could not open seccomp profile at", input)
		fmt.Println("[*] Creating new Profile")
	} else {
		jsonParser := json.NewDecoder(configFile)
		parseError := jsonParser.Decode(&SeccompProfile)
		fatalErrorCheck(parseError, "Error parsing Configuration File")
	}
	defer configFile.Close()

	parseSysCallFlag("kill", kill, &SeccompProfile)
	parseSysCallFlag("trap", trap, &SeccompProfile)
	parseSysCallFlag("errno", errno, &SeccompProfile)
	parseSysCallFlag("trace", trace, &SeccompProfile)
	parseSysCallFlag("allow", allow, &SeccompProfile)
	parseDefaultAction(defaultAction, &SeccompProfile)
	parseArchFlag(arch, &SeccompProfile)
	removeAction(remove, &SeccompProfile)

	b, marshallError := json.MarshalIndent(SeccompProfile, "", "    ")
	fatalErrorCheck(marshallError, "Error creating Seccomp Profile")

	newConfigFile, newConfigError := os.Create(name)
	fatalErrorCheck(newConfigError, "Error creating config file")
	_, writeError := newConfigFile.Write(b)
	fatalErrorCheck(writeError, "Error writing config to file")

}
