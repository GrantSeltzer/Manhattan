# Manhattan
A CLI tool for creating Docker seccomp json configurations

 TODO:
 - Write SeccompProfile back to a file
 - Functionality to remove actions completely
 - Allow arguments to be added for system calls
 - Go over code to make sure all errors are properly checked
 - Write unit tests
 - Need to check if syscall has specific arguments already when added
    - if It doesn't, the syscall action should be overwritten
    - If it does, a new one should be added
