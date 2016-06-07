package main

import "strings"

func parseFlagOpt(action string, arguments string) {

	var list []string
	if strings.Contains(arguments, ",") {
		list = strings.Split(arguments, ",")
	} else if strings.Contains(arguments, "/") {
		list = strings.Split(arguments, "/")
	}

	updateAction(action, list)
}

func updateAction(action string, list []string) {

}
