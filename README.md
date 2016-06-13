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

`manhattan -default kill`

`manhattan -default kill -remove clone`



##### To do
 - Allow arguments to be added for system calls

 - Write unit tests
 - Need to check if syscall has specific arguments already when added
    - if It doesn't, the syscall action should be overwritten
    - If it does, a new one should be added
