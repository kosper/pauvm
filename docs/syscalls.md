# Sycall Reference

Syscalls can be called using the **SYSCALL** instruction using the specified **opcode** of the syscall provided bellow.
Often, syscalls need multiple arguments, these arguments should be pushed onto the stack in order before calling the function. 
As such, the table bellow will specify the number of arguments and their types(_For now we only accept integer values this will change in the future_).

| Name | OpCode | # of args | Arg 1 | Arg 2 | Arg 3 | Arg 4 | Arg 5 | Description |
|------|------|------|------|------|------|------|------|------|
| `EXIT` | `0x00` | 1  | Exit value | - | - | - | - | Exits the program |
