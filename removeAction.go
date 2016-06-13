package main

import (
	"strings"

	"github.com/docker/engine-api/types"
)

func removeAction(arguments string, config *types.Seccomp) {
	var syscallsToRemove []string
	if strings.Contains(arguments, ",") {
		syscallsToRemove = strings.Split(arguments, ",")
	} else if strings.Contains(arguments, "/") {
		syscallsToRemove = strings.Split(arguments, "/")
	} else {
		syscallsToRemove = append(syscallsToRemove, arguments)
	}

	for _, syscall := range syscallsToRemove {
		for counter, syscallStruct := range config.Syscalls {
			if syscallStruct.Name == syscall {
				config.Syscalls = append(config.Syscalls[:counter], config.Syscalls[counter+1:]...)
			}
		}
	}
}
