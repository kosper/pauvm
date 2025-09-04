# PauVM
*A stack-based virtual machine and toolchain written in Go.*

---

## Overview
PauVM is a lightweight **stack-based virtual machine** with a custom bytecode format and additional tools for compiling-dissassembling.  
It includes:

- **pauvm**: Stack-Based Virtual Machine 
- **pauven**: Compiler that turns source files into bytecode.
- **paudiss**: Disassembler that turns bytecode files into readable code.

---

## Quickstart

### Build
```cmd
git clone https://github.com/kosper/pauvm.git
cd pauvm
cd build
build_all.bat
```

### Run Example(counter)
```cmd
cd bin
pauven.exe -f ../examples/counter.pv -o ../examples/counter.pau
pauvm.exe ../examples/counter.pau -trace
```
### Dissassemble Bytecode
```
cd bin
paudiss.exe ../examples/factorial.pau
```

---

### Example Code
Source(examples/factorial.pau)

```
#This program demonstrates the calculation of 5 factorial.
PUSH 5

CALL factorial
HALT

factorial:
  DUP
  JMPZ base_case
  DUP
  PUSH 1
  MINUS
  CALL factorial
  MUL
  RETURN

base_case:
  POP
  PUSH 1
  RETURN
```

Run
```cmd
cd bin
pauven.exe -f ../examples/factorial.pv -o ../examples/factorial.pau
pauvm.exe ../examples/factorial.pau -trace
```

---

### Testing 
Unit and integration tests are included. Run with:
```cmd
cd build
run_tests.exe

```

---

### Roadmap
- [ ] Documentation (IMPORTANT)
- [ ] Unit Tests.
- [ ] Linux build files.
- [ ] Implement Debugger
- [ ] Support including files in source code
- [ ] Handle arguments from console.
- [ ] Support aliases for values(macros)
- [ ] Native function calls/Syscalls
- [ ] OS exclusive error messages (e.g. usage right now is only for windows)
- [ ] CI
- [ ] Cleaner errors
- [ ] Profiling
- [ ] Add header when compiling to bytecode and design architecture of .pau file
- [ ] Flags for Compiler-VM-Dissassembler

---

### License
MIT License (see LICENSE).
