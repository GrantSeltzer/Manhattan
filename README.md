# Manhattan

A CLI tool/library for creating OCI seccomp json configurations.

#### [Why?](https://github.com/docker/docker/blob/master/docs/security/seccomp.md)

[![Go Report Card](https://goreportcard.com/badge/github.com/grantseltzer/manhattan)](https://goreportcard.com/report/github.com/grantseltzer/manhattan)

Manhattan is a tool used to generate the seccomp json file used by OCI containers to control the system calls available to processes running within a container.  The generated json files can be used by any OCI compliant runtime like runc and docker. You can pass them at the command line to docker like the following:

`docker run -it --security-opt seccomp:Manhattan.json fedora bash`

## Usage:

Arguments consist of all lower case names of [syscalls](http://man7.org/linux/man-pages/man2/syscalls.2.html). Multiple ones can be passed by using a `,` separated list.
Use any of the following flags to set [actions](https://www.kernel.org/doc/Documentation/prctl/seccomp_filter.txt) for specified syscalls:

`--kill` or `-k`

`--trap` or `-p`

`--errno` or `-e`

`--trace` or `-c`

`--allow` or `-a`

You can also specify [parameters](https://github.com/docker/engine-api/blob/master/types/seccomp.go#L51-L57) for rules to apply to. The syntax is as follows:

`manhattan --ACTION SYSCALL:INDEX:VALUE1:VALUE2:OP` OP must be any of the following:
`NE`, `LT`, `LE`, `EQ`, `GE`, `GT`, or `ME`.

`--remove` (`-r`) specifies syscalls that you would like to remove from the default configuration. Syscalls not specified will take on the default action.

`--default` (`-d`) specifies the default action for syscalls not explicitly specified.

`--arch` (`-l`)specifies supported [architectures](https://github.com/opencontainers/runc/blob/master/libcontainer/seccomp/config.go#L27-L44).

`--name` (`-n`) specifies the name of the output file. The default is the current timestamp in the current directory.

`--name-force` is the same as `--name` except it will overwrite an existing file if it's specified

### Library

Simply run `go get github.com/grantseltzer/manhattan/oci-seccomp-gen` and import it in your go project.

Documentation for use as a library coming soon.


## Example usages:
`manhattan --kill accept --name ~/jsonfiles/SeccompConfig`

`manhattan --input foo.bar --name-force foo.bar --kill clone:0:1:2:NE,getcwd`

`manhattan --kill=accept` , `manhattan --kill:accept` and `manattan --kill accept` are all equivalent

`manhattan --errno write,read --allow fstat`

`manhattan --remove clone`

`manhattan --default kill --remove clone`

`manhattan --trace clone:1:2:3:GT`

`manhattan --kill clone:1:2:3:ME,getcwd:1:2:3:GE`

`manhattan --arch mips,mips64,amd64`
