package pauvm

import (
	"testing"
)

func TestTraceFlag(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: true,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	trace := pauVM.trace

	if trace == false {
		t.Error("Trace flag was not set correctly for some reason.")
	}
}

func TestStackSize(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: 1024,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	stacksize := len(pauVM.stack)

	if stacksize != 1024 {
		t.Errorf("Expected stacksize of 1024, instead is %d.", stacksize)
	}
}

func TestCallstackSize(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: 64,
	}

	var pauVM *VM = InitVM(&flags)

	callstacksize := len(pauVM.frames)

	if callstacksize != 64 {
		t.Errorf("Expected callstack size of 64 , instead is %d.", callstacksize)
	}
}
