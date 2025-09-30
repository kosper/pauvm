package pauvm

import (
	"github.com/kosper/pauvm/pkg/isa"
)

const maxInstructions = 256
const safeInstructionCount = 256
const defaultStackSize = 256
const defaultCallstackSize = 256
const defaultLocalStorageSize = 64

type executionErrorType byte

//Note: Maybe i need to delete these so i don't clutter the memory with unimportant stuff.
//Note: After these are used, program exits anyway.
const (
	ERROR_STACK_OVERFLOW executionErrorType = iota
	ERROR_STACK_UNDERFLOW
	ERROR_DIV_BY_ZERO
	ERROR_PROGRAM_NOT_HALTED
	ERROR_ILLEGAL_JUMP
	ERROR_UKNOWN_INSTRUCTION
)

var errorToString = map[executionErrorType]string {
	ERROR_STACK_OVERFLOW: "Stack Overflow Error!",
	ERROR_STACK_UNDERFLOW: "Stack Underflow Error!",
	ERROR_DIV_BY_ZERO: "Division By Zero Error!",
	ERROR_PROGRAM_NOT_HALTED: "Program Not Halted Error!",
	ERROR_ILLEGAL_JUMP: "Illegal Jump Error!",

	//Note: Keep this for the immidiate mode, but note compiled files will never have uknown instrucitons.
	ERROR_UKNOWN_INSTRUCTION: "Uknown Instruction Error!",
}

type Instruction struct {
	inst isa.InstructionType
	value int32
}

//Note: Function frame
type Frame struct {
	returnIp int32
	locals [defaultLocalStorageSize]int32
}

//TODO: maxInstructions.
type VMFlags struct {
	Filename string
	trace bool
	stackSize int32
	callstackSize int32
}

type VM struct {
	program [maxInstructions]Instruction
	ip int32
	totalInstructions int32

	stack []int32
	sp int32
	
	frames []Frame	
	fp int32

	VMFlags
}
