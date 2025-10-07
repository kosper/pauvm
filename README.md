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

Or

```cmd
git clone https://github.com/kosper/pauvm.git
cd pauvm
go build ./...
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

main:
  PUSH 5
  
  CALL factorial
  HALT

factorial:
  DUP

  #This is not a function call, rather a way to demonstrate jumps.
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

  #Returns to main(Remember, not a function call).
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
Unit tests are included. Run with:
```cmd
cd build
run_tests.exe
```

Or

```cmd
go test ./...
```

### Build And Running A Docker Image

```cmd
cd pauvm
docker build -t pauvm .
```

After building the image, you can execute, compile or dissassemble bytecode by mounting the examples directory.

```cmd
docker run --rm -v ./examples:/examples pauvm pauvm examples/counter/counter.pau -trace
```

---

### Roadmap
- [X] Documentation (IMPORTANT)
- [X] Unit Tests.
- [X] Linux build files(NOTE: exist but not commited yet).
- [ ] Implement Debugger
- [X] Support including files in source code
- [X] Handle arguments from console(stack size, callstack size etc).
- [X] Support aliases for values(macros)
- [X] Native function calls/Syscalls support
- [X] Print line on error(Compiler)
- [X] OS exclusive error messages (e.g. usage right now is only for windows)
- [X] CI
- [X] Cleaner errors
- [ ] Profiling
- [X] Help flag 
- [X] Add header when compiling to bytecode and design architecture of .pau file
- [X] Flags for Compiler-VM-Dissassembler
- [ ] Data types like int64, int8, floats, doubles etc(NaN Boxing or Not).
- [ ] Strings
- [ ] Constant Pool

---

### License
MIT License (see LICENSE).
