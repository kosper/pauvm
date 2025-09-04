package pauvm

import (
	"testing"
	. "github.com/kosper/pauvm/pkg/isa"
)

//TODO: Better testing, test all instructions, test edge cases.
func TestPUSH(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 {
		t.Errorf("Expected 5 got %d", result)
	}
}

func TestADD(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_ADD, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 7 {
		t.Errorf("Expected 7 got %d", result)
	}
}

func TestMINUS(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_MINUS, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 3 {
		t.Errorf("Expected 3 got %d", result)
	}
}

func TestMUL(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_MUL, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 10 { 
		t.Errorf("Expected 10 got %d", result)
	}
}

func TestDIV(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_DIV, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 2 { 
		t.Errorf("Expected 2 got %d", result)
	}
}

func TestEQ(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_EQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}
}

func TestNEQ(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_NEQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}
}

func TestLS(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_LS, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}
}

func TestGR(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_GR, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}
}

func TestGREQ(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_GREQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}
}

func TestLSEQ(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_LSEQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}
}

func TestDUP(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_DUP, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 { 
		t.Errorf("Expected 5 got %d", result)
	}
}

func TestSWAP(t *testing.T) {
	var flags VMFlags = VMFlags {
		Trace: false,
	}

	var pauVM *VM = InitVM(&flags)

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_SWAP, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 { 
		t.Errorf("Expected 5 got %d", result)
	}
}
