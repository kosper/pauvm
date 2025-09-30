package pauvm

import (
	"errors"
	"os"
	"fmt"

	"github.com/kosper/pauvm/pkg/isa"
)

func (pauVM *VM) push() error {
	if pauVM.sp >= pauVM.stackSize {
		return errors.New(errorToString[ERROR_STACK_OVERFLOW])
	}
	
	pauVM.sp++
	var index int32 = pauVM.sp
	var ip int32 = pauVM.ip

	pauVM.stack[index] = pauVM.program[ip].value

	pauVM.ip++

	return nil;
}

func (pauVM *VM) add() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--
	pauVM.stack[pauVM.sp] = (value1 + value2)

	pauVM.ip++

	return nil;
}

func (pauVM *VM) minus() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--
	pauVM.stack[pauVM.sp] = (value2 - value1)

	pauVM.ip++

	return nil;
}

func (pauVM *VM) mul() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--
	pauVM.stack[pauVM.sp] = (value1 * value2)

	pauVM.ip++

	return nil;
}

func (pauVM *VM) div() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	if value1 == 0 {
		return errors.New(errorToString[ERROR_DIV_BY_ZERO])
	}
	
	pauVM.sp--
	pauVM.stack[pauVM.sp] = (value2 / value1)

	pauVM.ip++

	return nil;
}

func (pauVM *VM) mod() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--
	pauVM.stack[pauVM.sp] = (value2 % value1)

	pauVM.ip++

	return nil;
}

func (pauVM *VM) eq() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--

	if value1 == value2 {
		pauVM.stack[pauVM.sp] = 0
	} else {
		pauVM.stack[pauVM.sp] = 1
	}

	pauVM.ip++

	return nil;
}

func (pauVM *VM) neq() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--

	if value1 != value2 {
		pauVM.stack[pauVM.sp] = 0
	} else {
		pauVM.stack[pauVM.sp] = 1
	}

	pauVM.ip++

	return nil;
}

func (pauVM *VM) ls() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--

	if value1 < value2 {
		pauVM.stack[pauVM.sp] = 0
	} else {
		pauVM.stack[pauVM.sp] = 1
	}

	pauVM.ip++

	return nil;
}

func (pauVM *VM) gr() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--

	if value1 > value2 {
		pauVM.stack[pauVM.sp] = 0
	} else {
		pauVM.stack[pauVM.sp] = 1
	}

	pauVM.ip++

	return nil;
}

func (pauVM *VM) greq() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--

	if value1 >= value2 {
		pauVM.stack[pauVM.sp] = 0
	} else {
		pauVM.stack[pauVM.sp] = 1
	}

	pauVM.ip++

	return nil;
}

func (pauVM *VM) lseq() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.sp--

	if value1 <= value2 {
		pauVM.stack[pauVM.sp] = 0
	} else {
		pauVM.stack[pauVM.sp] = 1
	}

	pauVM.ip++

	return nil;
}

func (pauVM *VM) dup() error {
	if pauVM.sp < 0 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value int32 = pauVM.stack[pauVM.sp]

	pauVM.sp++
	pauVM.stack[pauVM.sp] = value

	pauVM.ip++

	return nil;
}

//TODO: Check if jump is illegal.
func (pauVM *VM) jmp() error {
	var value int32 = pauVM.program[pauVM.ip].value

	if value >= pauVM.totalInstructions {
		return errors.New(errorToString[ERROR_ILLEGAL_JUMP])
	}

	pauVM.ip = value;

	return nil
}

func (pauVM *VM) jmpz() error {
	if pauVM.sp < 0 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}

	var value int32 = pauVM.stack[pauVM.sp]
	pauVM.sp--;

	if value == 0 {
		var address = pauVM.program[pauVM.ip].value;

		if address >= pauVM.totalInstructions {
			return errors.New(errorToString[ERROR_ILLEGAL_JUMP])
		}

		pauVM.ip = address;
	} else {
		pauVM.ip++
	}

	return nil;
}

func (pauVM *VM) swap() error {
	if pauVM.sp < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.sp]
	var value2 int32 = pauVM.stack[pauVM.sp - 1]

	pauVM.stack[pauVM.sp - 1] = value1
	pauVM.stack[pauVM.sp] = value2

	pauVM.ip++

	return nil;
}

func (pauVM *VM)store() error {
	if pauVM.sp < 0 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}

	var value int32 = pauVM.stack[pauVM.sp]
	var index int32 = pauVM.program[pauVM.ip].value;

	pauVM.sp--;

	if index < 0 || index > defaultLocalStorageSize {
		return errors.New("Local storage index out of bounds error")
	}

	var fp int32 = pauVM.fp
	pauVM.frames[fp].locals[index] = value

	pauVM.ip++

	return nil;
}

func (pauVM *VM)load() error {
	var index int32 = pauVM.program[pauVM.ip].value;

	if index < 0 || index > defaultLocalStorageSize {
		return errors.New("Local storage index out of bounds error")
	}

	var fp int32 = pauVM.fp
	var value int32 = pauVM.frames[fp].locals[index] 

	pauVM.sp++
	pauVM.stack[pauVM.sp] = value

	pauVM.ip++

	return nil;
}

func (pauVM *VM)call() error {
	var fp int32 = pauVM.fp

	if fp >= pauVM.callstackSize {
		return errors.New("Callstack overflow.")
	}

	pauVM.fp++
	pauVM.frames[pauVM.fp].returnIp = pauVM.ip + 1

	var label int32 = pauVM.program[pauVM.ip].value

	if label >= pauVM.totalInstructions {
		return errors.New(errorToString[ERROR_ILLEGAL_JUMP])
	}

	pauVM.ip = label
	
	return nil;
}

func (pauVM *VM)syscall() error {
	var call int32 = pauVM.program[pauVM.ip].value;
	
	switch isa.Syscall(call) {
		case isa.SYSCALL_EXIT:
			if pauVM.sp < 0 {
				return errors.New("Stack underflow!")
			}

			var value int32 = pauVM.stack[pauVM.sp]
			pauVM.sp--;		
			
			os.Exit(int(value))

			break;
		default:
			var ferror string = fmt.Sprintf("Syscall %d does not exist.", call)
			return errors.New(ferror)
	}
	return nil;
}

func (pauVM *VM)ret() error {
	var fp int32 = pauVM.fp

	if fp <= 0 {
		return errors.New("Callstack underflow.")
	}

	pauVM.ip = pauVM.frames[fp].returnIp
	pauVM.frames[fp].returnIp = 0 

	pauVM.fp--


	return nil;
}

func (pauVM *VM) pop() error {
	if pauVM.sp < 0 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}

	pauVM.sp--
	pauVM.ip++

	return nil
}
