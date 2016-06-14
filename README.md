# Manhattan
A CLI tool for creating Docker seccomp json configurations. [Why?](https://github.com/docker/docker/blob/master/docs/security/seccomp.md)

[![Go Report Card](https://goreportcard.com/badge/github.com/grantseltzer/manhattan)](https://goreportcard.com/report/github.com/grantseltzer/manhattan)

## Usage:

Use any of the following flags to set actions for specified syscalls: (Specifying arguments coming soon!)

`-kill`
`-trap`
`-errno`
`-trace`
`-allow`

Arguments consist of all lower case names of syscalls. Multiple ones can be passed by using a `,` or `/` separated list

You can also specify parameters for rules to apply to. The syntax is as follows:
`manhattan -ACTION SYSCALL:INDEX:VALUE1:VALUE2:OP` OP must be any of the following:
`NE`, `LT`, `LE`, `EQ`, `GE`, `GT`, or `ME`. For an explanation see
[here](https://github.com/docker/engine-api/blob/master/types/seccomp.go#L51-L57)

`-remove` specifies syscalls that you would like to remove from the default configuration. Syscalls not specified will take on the default action.

`-default` specifies the default action for syscalls not explicitly specified.

You can find an explanation of seccomp actions [here](https://www.kernel.org/doc/Documentation/prctl/seccomp_filter.txt)

You can find a list of syscalls [here](http://man7.org/linux/man-pages/man2/syscalls.2.html)


##### Output will be in the form of a docker compliant JSON file.

`-location` specifies the location of the ouput file. The default is the home directory.

`-name` specifies the name of the output file. The default is the current timestamp.

## Example usages:
`manhattan -kill accept -location ~/jsonfiles -name SeccompConfig`

`manhattan -kill=accept` , `manhattan -kill:accept` and `manattan -kill accept` are all equivalent

`manhattan -errno write,read -allow fstat`

`manhattan -default kill -remove clone`

`manhattan -trace clone:1:2:3:GT`


##### To do
 - man page
 - -- for full word flags, - for single letter flags
 - Architecture flag
 - input default profile option
 - if neither can be found, build one from scratch
 - makefile
 - Write unit tests
