package main

import (
	"encoding/json"
	"fmt"
	"os"

	seccomp "github.com/grantseltzer/Manhattan/ociseccompgen"
	types "github.com/opencontainers/runtime-spec/specs-go"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	defaultSeccompProfile = "/etc/sysconfig/manhattan.json"
	version               = "0.0.1"
)

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
		nameforce     string
	)
	app := cli.NewApp()
	app.Name = "manhattan"
	app.Usage = "Create seccomp json configurations for use with OCI or Docker"
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
			Name:        "trap",
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
			Name:        "trace",
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
			Value:       "allow",
			Usage:       "Set the default action for syscalls not specified",
			Destination: &defaultAction,
		},
		cli.StringFlag{
			Name:        "arch",
			Value:       "amd64,x86,x32",
			Usage:       "Set supported architectures",
			Destination: &arch,
		},
		cli.StringFlag{
			Name:        "name, n",
			Value:       seccomp.DefaultFullPath(),
			Usage:       "Set name of output file",
			Destination: &name,
		},
		cli.StringFlag{
			Name:        "name-force",
			Value:       "not-specified",
			Usage:       "Set name of output file, force write",
			Destination: &nameforce,
		},
	}
	app.Version = version
	app.EnableBashCompletion = true
	app.Action = func(c *cli.Context) error {
		return nil
	}

	app.Run(os.Args)

	var SeccompProfile types.Seccomp

	configFile, err := os.Open(input)
	if err != nil {
		fmt.Println("[*] Could not open seccomp profile at", input)
		fmt.Println("[*] Creating new Profile")
	} else {
		jsonParser := json.NewDecoder(configFile)
		err = jsonParser.Decode(&SeccompProfile)
		if err != nil {
			logrus.Fatal("Error parsing Configuration File", err)
		}
		defer configFile.Close()
	}

	err = seccomp.DefaultAction(defaultAction, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing default action argument", err)
	}

	err = seccomp.ParseSyscallFlag("kill", kill, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing kill argument", err)
	}

	err = seccomp.ParseSyscallFlag("trap", trap, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing trap argument", err)
	}

	err = seccomp.ParseSyscallFlag("errno", errno, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing errno argument", err)
	}

	err = seccomp.ParseSyscallFlag("trace", trace, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing trace argument", err)
	}

	err = seccomp.ParseSyscallFlag("allow", allow, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing allow argument", err)
	}

	err = seccomp.ParseArchitectureFlag(arch, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing architecture agument", err)
	}

	err = seccomp.RemoveAction(remove, &SeccompProfile)
	if err != nil {
		logrus.Fatal("Error parsing remove action argument", err)
	}

	b, err := json.MarshalIndent(SeccompProfile, "", "    ")
	if err != nil {
		logrus.Fatal("Error creating Seccomp Profile", err)
	}

	if _, erro := os.Stat(name); erro == nil && nameforce == "not-specified" {
		logrus.Fatal("File destination already exists. Use --name-force to overwrite. ", name, erro)
	}

	if nameforce != "not-specified" {
		name = nameforce
	}

	newConfigFile, err := os.Create(name)
	if err != nil {
		logrus.Fatal("Error creating config file", err)
	}
	if _, err := newConfigFile.Write(b); err != nil {
		logrus.Fatal("Error writing config to file", err)
	}
}

func autocomplete(c *cli.Context) {
	tasks := []string{"kill", "trap", "errno", "trace", "allow", "remove",
		"default", "arch", "name"}

	// This will complete if no args are passed
	if c.NArg() > 0 {
		return
	}
	for _, t := range tasks {
		fmt.Println(t)
	}
}
