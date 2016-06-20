package main

import (
	"fmt"

	"github.com/urfave/cli"
)

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
