package main

import "flag"

func main() {

	kill := flag.String("kill", "default", "Respond to system call with KILL")
	trap := flag.String("trap", "default", "Respond to system call with TRAP")
	errno := flag.String("errno", "default", "Respond to system call with ERRNO")
	trace := flag.String("trace", "default", "Respond to system call with TRACE")
	allow := flag.String("allow", "default", "Respond to system call with ALLOW")

	flag.Parse()

	parseFlagOpt("kill", *kill)
	parseFlagOpt("trap", *trap)
	parseFlagOpt("errno", *errno)
	parseFlagOpt("trace", *trace)
	parseFlagOpt("allow", *allow)

}
