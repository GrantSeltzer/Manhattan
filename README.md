# Manhattan
A CLI tool for creating Docker seccomp json configurations. [Why?](https://github.com/docker/docker/blob/master/docs/security/seccomp.md)

## Usage:

Use any of the following flags to set actions for specified syscalls: (Specifying arguments coming soon!)

`-kill`
`-trap`
`-errno`
`-trace`
`-allow`

Arguments consist of all lower case names of syscalls. Multiple ones can be passed by using a `,` or `/` separated list

`-remove` COMING SOON

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



 #### TODO:
 - Functionality to remove actions completely
 - Allow arguments to be added for system calls
 - Go over code to make sure all errors are properly checked
 - Write unit tests
 - Need to check if syscall has specific arguments already when added
    - if It doesn't, the syscall action should be overwritten
    - If it does, a new one should be added
