package pauvm

//TODO: Replace prints with log

import (
	"fmt"
	"errors"

	"os"
	"io"
	"encoding/binary"

	"github.com/kosper/pauvm/pkg/isa"
)

func InitVM(flags *VMFlags) *VM {
	//NOTE: These values are already initialized by golang to 0.
	var pauVM VM = VM{ 
		ip: 0, 
		totalInstructions: 0, 
		stackIndex: 0, 
		fp: 0,
	}
	
	pauVM.Trace = flags.Trace

	for i := range maxInstructions {
		pauVM.program[i].inst = isa.INST_NONE
	}

	return &pauVM
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

	for i := range pauVM.stackIndex {
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

			case isa.INST_RETURN:
				err = pauVM.ret()
				break

			case isa.INST_HALT:
				return nil
			
			case isa.INST_NONE:
				continue

			default:
				return errors.New(errorToString[ERROR_UKNOWN_INSTRUCTION]);
		}

		if err != nil {
			return err
		}

		if pauVM.Trace == true {
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
