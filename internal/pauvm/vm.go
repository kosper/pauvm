package pauvm

import (
	"fmt"
	"errors"

	"os"
	"runtime"
	"strconv"
	"io"
	"encoding/binary"

	"github.com/kosper/pauvm/pkg/isa"
	"github.com/kosper/pauvm/pkg/utils"

)

func printUsage() {
	switch runtime.GOOS {
		case "windows": {
			var UsageWindows string = "Usage: pauvm.exe <bytecodeFile> [-flags]"
			fmt.Println(UsageWindows)

			return
		}

		case "linux": {
			var UsageLinux string = "Usage: ./pauvm <bytecodeFile> [-flags]"
			fmt.Println(UsageLinux)

			return
		}

		default: {
			var UsageLinux string = "Usage: ./pauvm <bytecodeFile> [-flags]"
			fmt.Println(UsageLinux)

			return
		}
	}
}

func printHelp() {
	var helpString string = 
`PauVM is a stack based virtual machine that executes Pau bytecode.
Usage:
	pauvm <bytecode file> [flags]	

The flags are:
  -help                  Prints the help menu.
  -stacksize <size>      Changes the size of the stack to the specified size.
  -trace                 Prints the stack everytime an instruction is executed.
  -callstacksize <size>  Changes the size of the callstack to the specified size.
`
		
	fmt.Println(helpString)
}

func HandleConsoleArgs() (*VMFlags, error) {
	var args []string = os.Args[1:]
	var argsLen = len(args)

	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	if argsLen < 1 {
		printUsage()
		os.Exit(0)
	}

	//Note: Check if first flag is help, if not continue to other flags.
	if args[0] == "-help" {
		printHelp()
		os.Exit(0)
	}

	flags.Filename = args[0]

	if err := utils.IsFileExtension(flags.Filename, ".pau"); err != nil {
		return &flags, err
	}

	//Note: First argument is filename or help flag which is handled.
	for i := 1; i < argsLen; i++ {
		//Note: Handle flags.
		switch args[i] {
			case "-trace":
				flags.trace = true
				continue
			case "-stacksize":
				if i + 1 <= len(args) {
					var ferror string = "Error: Stacksize was not provided" 
					return &flags, errors.New(ferror)
				} 

				value, err := strconv.ParseInt(args[i + 1], 10, 32)

				if err != nil {
					return &flags, err
				}

				flags.stackSize = int32(value)

				i++

				continue; 
			case "-callstacksize":
				if i + 1 <= len(args) {
					var ferror string = "Error: Stacksize was not provided" 
					return &flags, errors.New(ferror)
				}

				value, err := strconv.ParseInt(args[i + 1], 10, 32)

				if err != nil {
					return &flags, err
				}

				flags.callstackSize= int32(value)

				i++	

				continue;
			default:
				var ferror = fmt.Sprintf("Error: %s flag does not exist.", args[i])
				return &flags, errors.New(ferror)
		}
		
	}

	return &flags, nil
}

func InitVM(flags *VMFlags) *VM {
	//Note: These values are already initialized by golang to 0.
	var pauVM VM = VM{ 
		ip: 0, 
		totalInstructions: 0, 
		sp: -1, 
		fp: 0,
	}
	
	pauVM.trace = flags.trace
	pauVM.stackSize = flags.stackSize
	pauVM.callstackSize = flags.callstackSize

	pauVM.stack = make([]int32, flags.stackSize)
	pauVM.frames = make([]Frame, flags.callstackSize)

	for i := range maxInstructions {
		pauVM.program[i].inst = isa.INST_NONE
	}

	return &pauVM
}

//Note: Resets the VM.
func (pauVM *VM)Reset() {
	pauVM.ip = 0
	pauVM.totalInstructions = 0
	pauVM.sp = -1 
	pauVM.fp = 0 

	for i := range maxInstructions {
		pauVM.program[i].inst = isa.INST_NONE
	}
}

func (pauVM *VM) AddInstruction(inst isa.InstructionType, value int32) error {
	if pauVM.totalInstructions >= maxInstructions {
		return errors.New("Too many instructions!");
	}

	var index int32 = pauVM.totalInstructions;

	pauVM.program[index].inst = inst;
	pauVM.program[index].value = value;

	pauVM.totalInstructions++;

	return nil
}

func (pauVM *VM) PrintProgram() {
	fmt.Printf("Program:\n");

	for i := range pauVM.totalInstructions {
		var instName string = isa.InstToString[pauVM.program[i].inst]
		var value int32 = pauVM.program[i].value

		fmt.Printf("Instruction(%v): %v, Value: %v\n", i, instName, value);
	}
	
	fmt.Printf("\n");
	
	return
}

func (pauVM *VM) PrintStack() {
	fmt.Printf("Stack:\n");

	for i := range pauVM.sp + 1 {
		fmt.Printf("Address(%v): %v\n", i, pauVM.stack[i]);
	}

	fmt.Printf("\n");

	return
}

//TODO: Array with function pointers to operation functions instead of switching.
func (pauVM *VM) ExecuteProgram() error {
	for _ = range safeInstructionCount {
		var inst = pauVM.program[pauVM.ip].inst
		var err error = nil
		var val int32 = pauVM.program[pauVM.ip].value;	

		switch inst {
			case isa.INST_PUSH:
				err = pauVM.push() 
				break 

			case isa.INST_POP:
				err = pauVM.pop() 
				break 

			case isa.INST_ADD:
				err = pauVM.add() 
				break 

			case isa.INST_MINUS:
				err = pauVM.minus() 
				break 

			case isa.INST_MUL:
				err = pauVM.mul() 
				break 

			case isa.INST_DIV:
				err = pauVM.div() 
				break 

			case isa.INST_MOD:
				err = pauVM.mod() 
				
			case isa.INST_EQ:
				err = pauVM.eq() 
				break 

			case isa.INST_NEQ:
				err = pauVM.neq() 
				break 

			case isa.INST_LS:
				err = pauVM.ls() 
				break 

			case isa.INST_GR:
				err = pauVM.gr() 
				break 
			
			case isa.INST_GREQ:
				err = pauVM.greq()
				break 

			case isa.INST_LSEQ:
				err = pauVM.lseq()
				break 

			case isa.INST_DUP:
				err = pauVM.dup()
				break

			case isa.INST_JMP:
				err = pauVM.jmp()
				break

			case isa.INST_JMPZ:
				err = pauVM.jmpz()
				break

			case isa.INST_SWAP:
				err = pauVM.swap()
				break

			case isa.INST_STORE:
				err = pauVM.store()
				break

			case isa.INST_LOAD:
				err = pauVM.load()
				break

			case isa.INST_CALL:
				err = pauVM.call()
				break

			case isa.INST_SYSCALL:
				err = pauVM.syscall()
				break

			case isa.INST_RETURN:
				err = pauVM.ret()
				break

			case isa.INST_HALT:
				if pauVM.trace == true {
					fmt.Printf("Inst HALT\n")
				}
				return nil
			
			case isa.INST_NONE:
				continue

			default:
				return errors.New(errorToString[ERROR_UKNOWN_INSTRUCTION]);
		}

		if err != nil {
			return err
		}

		if pauVM.trace == true {
			fmt.Printf("Inst %v, Value %v\n", isa.InstToString[inst], val);
			pauVM.PrintStack()
		}
	}

	return errors.New(errorToString[ERROR_PROGRAM_NOT_HALTED])
}

func (pauVM *VM) LoadProgramFromFile(filename string) error {
	var file, err = os.Open(filename)
	defer file.Close()

	if err != nil {
		return err
	}

	//Note: Read header.
	header, err := utils.ReadHeader(file)

	if err != nil {
		return err
	}

	if err = utils.CheckMagicNumber(header); err != nil {
		return err
	}

	var instruction []byte = make([]byte, 1) 

	for {
		_, err := io.ReadFull(file, instruction)

		if err == io.EOF { break }
		if err != nil { return err }

		opcode := isa.InstructionType(instruction[0])

		var value int32 = 0

		if isa.InstHasOperand[opcode] {
			if err := binary.Read(file, binary.LittleEndian, &value); err != nil {
				return err
			}
		}

		if err := pauVM.AddInstruction(opcode, value); err != nil {
			return err
		}
	}

	return nil
}
