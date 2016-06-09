package main

import (
	"strings"
	"time"
)

//returns current time and date as a string without any whitespace
func defaultTime() string {
	return strings.Replace(time.Now().String(), " ", "", -1)
}
