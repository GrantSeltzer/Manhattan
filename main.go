package main

import (
	"encoding/json"
	"fmt"
	"os"

	seccomp "github.com/grantseltzer/Manhattan/ociseccompgen"
	types "github.com/opencontainers/runc/libcontainer/configs"

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
			Value:       "errno",
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
		if jsonParser.Decode(&SeccompProfile) != nil {
			logrus.Fatal("Error parsing Configuration File")
		}
		defer configFile.Close()
	}

	if seccomp.ParseSyscallFlag("kill", kill, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing kill argument")
	}
	if seccomp.ParseSyscallFlag("trap", trap, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing trap argument")
	}
	if seccomp.ParseSyscallFlag("errno", errno, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing errno argument")
	}
	if seccomp.ParseSyscallFlag("trace", trace, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing trace argument")
	}
	if seccomp.ParseSyscallFlag("allow", allow, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing allow argument")
	}
	if seccomp.DefaultAction(defaultAction, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing default action argument")
	}
	if seccomp.ParseArchitectureFlag(arch, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing architecture agument")
	}
	if seccomp.RemoveAction(remove, &SeccompProfile) != nil {
		logrus.Fatal("Error parsing remove action argument")
	}

	b, err := json.MarshalIndent(SeccompProfile, "", "    ")
	if err != nil {
		logrus.Fatal("Error creating Seccomp Profile")
	}

	newConfigFile, err := os.Create(name)
	if err != nil {
		logrus.Fatal("Error creating config file")
	}

	if _, err := newConfigFile.Write(b); err != nil {
		logrus.Fatal("Error writing config to file")
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
