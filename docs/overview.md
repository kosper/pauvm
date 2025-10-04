# Architecture Overview

## Pauven
- **Preprocessor**: Parses the main file, handles **macros**, **includes**, **functions** and **labels**.
- **Compiler**: Parses the preprocessed data, translates data to bytecode and writes to file.

## Pauvm
- **VM**: Parses a compiled _pau_ files, loads the instructions into memory and executes them.

## Paudiss
- **Dissassembler**: Parses a compiled _pau_ file and translates opcodes into human-readable instructions.

---

## Pau File Memory Layout
[MagicNumber-4 bytes][Version-3 bytes][call main][Program]

## PV File
### Rules
- Every *pv* file should include a main function. This is the first function that is called by the VM.
- Comments can be written using hashtag(#) everywhere, exept in the same line as **macro defines**,**includes** and **labels** or **function definitions**. Comments are ignored by the compiler.
- _pv_ files can have empty lines, they are ignored by the compiler.
- Identation does not matter.
- If at the end of the line the (**:**) symbol is detected, eveything in that line is considared a **label-function** name.
- If at the start of the line the (**!**) symbol is detected. Everything after that symbol is handled by the preprocessor.
- Macros can be defined by using `!def <Macro> <Definition>`, and can be used recursively(Meaning a macro can define another macro).
- Files can be included be using `!include <filepath>`.
- Preprocessor keeps track of included files so do not worry of including the same file twice.
- See the _/examples/_ folder for more example projects.
