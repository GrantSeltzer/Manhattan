package main

import (
	"encoding/json"
	"fmt"
	"os"

	parse "github.com/grantseltzer/Manhattan/parse"

	log "github.com/Sirupsen/logrus"
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
			Value:       parse.DefaultFullPath(),
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
		if parseError != nil {
			log.Fatal("Error parsing Configuration File")
		}
	}
	defer configFile.Close()

	parse.SysCallFlag("kill", kill, &SeccompProfile)
	parse.SysCallFlag("trap", trap, &SeccompProfile)
	parse.SysCallFlag("errno", errno, &SeccompProfile)
	parse.SysCallFlag("trace", trace, &SeccompProfile)
	parse.SysCallFlag("allow", allow, &SeccompProfile)
	parse.DefaultAction(defaultAction, &SeccompProfile)
	parse.ArchFlag(arch, &SeccompProfile)
	parse.RemoveAction(remove, &SeccompProfile)

	b, marshallError := json.MarshalIndent(SeccompProfile, "", "    ")
	if marshallError != nil {
		log.Fatal("Error creating Seccomp Profile")
	}

	newConfigFile, newConfigError := os.Create(name)
	if newConfigError != nil {
		log.Fatal("Error creating config file")
	}

	_, writeError := newConfigFile.Write(b)
	if writeError != nil {
		log.Fatal("Error writing config to file")
	}

}
