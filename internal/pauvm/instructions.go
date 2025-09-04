package pauvm

import (
	"errors"
)

//TODO: Right now the sp does not show the head of the stack, but rather the one next
//I should really change it to point to the head(maybe initialize it to -1???).

func (pauVM *VM) push() error {
	if pauVM.stackIndex >= stackSize {
		return errors.New(errorToString[ERROR_STACK_OVERFLOW])
	}
	
	var index int32 = pauVM.stackIndex
	var ip int32 = pauVM.ip

	pauVM.stack[index] = pauVM.program[ip].value

	pauVM.ip++
	pauVM.stackIndex++

	return nil;
}

func (pauVM *VM) add() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	pauVM.stack[pauVM.stackIndex - 2] = (value1 + value2)

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) minus() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	pauVM.stack[pauVM.stackIndex - 2] = (value2 - value1)

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) mul() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	pauVM.stack[pauVM.stackIndex - 2] = (value1 * value2)

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) div() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	if value1 == 0 {
		return errors.New(errorToString[ERROR_DIV_BY_ZERO])
	}
	
	pauVM.stack[pauVM.stackIndex - 2] = (value2 / value1)

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) eq() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	if value1 == value2 {
		pauVM.stack[pauVM.stackIndex - 2] = 0
	} else {
		pauVM.stack[pauVM.stackIndex - 2] = 1
	}

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) neq() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	if value1 != value2 {
		pauVM.stack[pauVM.stackIndex - 2] = 0
	} else {
		pauVM.stack[pauVM.stackIndex - 2] = 1
	}

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) ls() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	if value1 < value2 {
		pauVM.stack[pauVM.stackIndex - 2] = 0
	} else {
		pauVM.stack[pauVM.stackIndex - 2] = 1
	}

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) gr() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	if value1 > value2 {
		pauVM.stack[pauVM.stackIndex - 2] = 0
	} else {
		pauVM.stack[pauVM.stackIndex - 2] = 1
	}

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) greq() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	if value1 >= value2 {
		pauVM.stack[pauVM.stackIndex - 2] = 0
	} else {
		pauVM.stack[pauVM.stackIndex - 2] = 1
	}

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) lseq() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	if value1 <= value2 {
		pauVM.stack[pauVM.stackIndex - 2] = 0
	} else {
		pauVM.stack[pauVM.stackIndex - 2] = 1
	}

	pauVM.stackIndex--
	pauVM.ip++

	return nil;
}

func (pauVM *VM) dup() error {
	if pauVM.stackIndex < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value int32 = pauVM.stack[pauVM.stackIndex - 1]

	pauVM.stack[pauVM.stackIndex] = value

	pauVM.stackIndex++
	pauVM.ip++

	return nil;
}

func (pauVM *VM) jmp() error {
	pauVM.ip = pauVM.program[pauVM.ip].value;

	return nil
}

func (pauVM *VM) jmpz() error {
	if pauVM.stackIndex < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}

	pauVM.stackIndex--;
	var value int32 = pauVM.stack[pauVM.stackIndex]

	if value == 0 {
		pauVM.ip = pauVM.program[pauVM.ip].value;
	} else {
		pauVM.ip++
	}

	return nil;
}

func (pauVM *VM) swap() error {
	if pauVM.stackIndex < 2 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}
	
	var value1 int32 = pauVM.stack[pauVM.stackIndex - 1]
	var value2 int32 = pauVM.stack[pauVM.stackIndex - 2]

	pauVM.stack[pauVM.stackIndex - 2] = value1
	pauVM.stack[pauVM.stackIndex - 1] = value2

	pauVM.ip++

	return nil;
}

func (pauVM *VM)store() error {
	if pauVM.stackIndex < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}

	pauVM.stackIndex--;

	var value int32 = pauVM.stack[pauVM.stackIndex]
	var index int32 = pauVM.program[pauVM.ip].value;

	if index < 0 || index > localStorageSize {
		return errors.New("Local storage index out of bounds error")
	}

	var fp int32 = pauVM.fp
	pauVM.frames[fp].locals[index] = value

	pauVM.ip++

	return nil;
}

func (pauVM *VM)load() error {
	var index int32 = pauVM.program[pauVM.ip].value;

	if index < 0 || index > localStorageSize {
		return errors.New("Local storage index out of bounds error")
	}

	var fp int32 = pauVM.fp
	var value int32 = pauVM.frames[fp].locals[index] 

	pauVM.stack[pauVM.stackIndex] = value
	pauVM.stackIndex++

	pauVM.ip++

	return nil;
}

func (pauVM *VM)call() error {
	var fp int32 = pauVM.fp

	if fp >= frameStackSize {
		return errors.New("Framestack overflow.")
	}

	pauVM.fp++
	pauVM.frames[pauVM.fp].returnIp = pauVM.ip + 1

	var label int32 = pauVM.program[pauVM.ip].value

	pauVM.ip = label
	
	return nil;
}

func (pauVM *VM)ret() error {
	var fp int32 = pauVM.fp

	if fp <= 0 {
		return errors.New("Framestack underflow.")
	}

	pauVM.ip = pauVM.frames[fp].returnIp
	pauVM.frames[fp].returnIp = 0 

	pauVM.fp--


	return nil;
}

func (pauVM *VM) pop() error {
	if pauVM.stackIndex < 1 {
		return errors.New(errorToString[ERROR_STACK_UNDERFLOW])
	}

	pauVM.stackIndex--
	pauVM.ip++
	return nil
}
