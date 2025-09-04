package pauvm

import (
	"github.com/kosper/pauvm/pkg/isa"
)

const maxInstructions = 256
const safeInstructionCount = 256
const stackSize = 256
const frameStackSize = 256
const localStorageSize = 64

type executionErrorType byte

const (
	ERROR_STACK_OVERFLOW executionErrorType = iota
	ERROR_STACK_UNDERFLOW
	ERROR_DIV_BY_ZERO
	ERROR_PROGRAM_NOT_HALTED
	ERROR_UKNOWN_INSTRUCTION
)

//TODO: customize to accept formatting for lines, indices etc.
var errorToString = map[executionErrorType]string {
	ERROR_STACK_OVERFLOW: "Stack Overflow Error!",
	ERROR_STACK_UNDERFLOW: "Stack Underflow Error!",
	ERROR_DIV_BY_ZERO: "Division By Zero Error!",
	ERROR_PROGRAM_NOT_HALTED: "Program Not Halted Error!",
	ERROR_UKNOWN_INSTRUCTION: "Uknown Instruction Error!",
}

type Instruction struct {
	inst isa.InstructionType
	value int32
}

//NOTE: Function frame
type Frame struct {
	returnIp int32
	locals [localStorageSize]int32
}

type VMFlags struct {
	Trace bool
}

type VM struct {
	program [maxInstructions]Instruction
	ip int32
	totalInstructions int32

	stack [stackSize]int32
	sp int32
	
	frames [frameStackSize]Frame	
	fp int32

	VMFlags
}
