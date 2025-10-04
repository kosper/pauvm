# Bytecode Reference

| Name | OpCode | Has Arg | Description |
|------|--------|---------|-------------|
| `PUSH` | `0x00` | Yes | Pushes a value at the top of the stack. |
| `POP` | `0x01` | No | Removes a value from the top of the stack. |
| `ADD` | `0x02` | No | Adds the 2 values from the top of the stack and pushes the result onto the stack. |
| `MINUS` | `0x03` | No | Subtracts the 2 values from the top of the stack and pushes the result onto the stack. |
| `MUL` | `0x04` | No | Multiples the 2 values from the top of the stack and pushes the result onto the stack. |
| `DIV` | `0x05` | No | Divides the 2 values from the top of the stack and pushes the result onto the stack. |
| `MOD` | `0x06` | No | Calculates the remainder from the 2 values at the top of the stack, and pushes the result onto the stack. |
| `EQ` | `0x07` | No | Removes the last 2 values from the stack and pushes **0** if they are equal and **1** if they are not. |
| `NEQ` | `0x08` | No | Removes the last 2 values from the stack and pushes **0** if they are not equal and **1** if they are. |
| `LS` | `0x09` | No | Removes the last 2 values from the stack and pushes **0** if the first value is less than the second and **1** if not. |
| `GR` | `0x0A` | No | Removes the last 2 values from the stack and pushes **0** if the first value is greater than the second and **1** if not. |
| `GREQ` | `0x0B` | No | Removes the last 2 values from the stack and pushes **0** if the first value is greater or equal than the second and **1** if not. |
| `LSEQ` | `0x0C` | No | Removes the last 2 values from the stack and pushes **0** if the first value is less or equal than the second and **1** if not. |
| `DUP` | `0x0E` | No | Duplicates the value from the top of the stack. |
| `JMP` | `0x0F` | Yes | Jumps to the instruction specified in the argument. Argument can be a **label**, or an **index** to the instruction pointer. |
| `JMPZ` | `0x10` | Yes | Jumps to the instruction specified in the argument if the value at the top of the stack is **0**. Argument can be a **label**, or an **index** to the instruction pointer. |
| `SWAP` | `0x11` | No | Swaps the last 2 values at the top of the stack. |
| `STORE` | `0x12` | Yes | Removes and stores the last value at the top of the stack to the local memory **index** specified as an argument. |
| `LOAD` | `0x13` | Yes | Pushes onto the stack a value from the local memory of the function specified as an argument. |
| `CALL` | `0x14` | Yes | Calls the function specified as an argument. |
| `SYSCALL` | `0x15` | Yes | Calls a system function specified as an argument(_check syscalls.md_). |
| `RETURN` | `0x16` | No | Returns from a function. |
| `HALT` | `0x17` | No | Terminates the program. |
| `NONE` | `0x18` | No | Does nothing at all. |
