package pauvm

import (
	"testing"
	. "github.com/kosper/pauvm/pkg/isa"
)

func TestPUSH(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Test regular pushing onto the stack.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 {
		t.Errorf("Expected 5 got %d", result)
	}

	//Note: Push 2 after an operation. 
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, -2)
	pauVM.AddInstruction(INST_PUSH, -3)
	pauVM.AddInstruction(INST_ADD, 0)

	pauVM.AddInstruction(INST_PUSH, -42)

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != -42 {
		t.Errorf("Expected -42 got %d", result)
	}
}

func TestPOP(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Testing POP
	pauVM.AddInstruction(INST_PUSH, -22)
	pauVM.AddInstruction(INST_PUSH, 32)
	pauVM.AddInstruction(INST_POP, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != -22 {
		t.Errorf("Expected -22 got %d", result)
	}

	//Note: Verify that i can't POP on an empty stack.
	pauVM.Reset()
	pauVM.AddInstruction(INST_POP, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected error when popping on an emppty stack.")
	}
}

func TestADD(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic addition.
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

	//Note: Verify that addition canont happen we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 5)
	pauVM.AddInstruction(INST_ADD, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestMINUS(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic subtraction.
	pauVM.AddInstruction(INST_PUSH, 100)	
	pauVM.AddInstruction(INST_PUSH, 20)	
	pauVM.AddInstruction(INST_MINUS, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 80 {
		t.Errorf("Expected 80 got %d", result)
	}

	//Note: Verify that subtraction canont happen we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, -12)
	pauVM.AddInstruction(INST_MINUS, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestMUL(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify multiplication.
	pauVM.AddInstruction(INST_PUSH, 10)	
	pauVM.AddInstruction(INST_PUSH, 10)	
	pauVM.AddInstruction(INST_MUL, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 100 { 
		t.Errorf("Expected 100 got %d", result)
	}

	//Note: Verify that multiplication canont happen we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 10)
	pauVM.AddInstruction(INST_MUL, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestDIV(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic division.
	pauVM.AddInstruction(INST_PUSH, 10)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_DIV, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 { 
		t.Errorf("Expected 5 got %d", result)
	}

	//Note: Verify that division canont happen we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_DIV, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}

	//Note: Verify that we cannot divide by 0.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_PUSH, 0)
	pauVM.AddInstruction(INST_DIV, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected DivisionByZeroError")
	}
}

func TestMOD(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Test basic modulo.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_MOD, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}

	//Note: Verify that modulo canont happen if we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_MOD, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestEQ(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify that these 2 elements are not equal.
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

	//Note: Verify that these 2 elements are equal.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, -32)	
	pauVM.AddInstruction(INST_PUSH, -32)	
	pauVM.AddInstruction(INST_EQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}

	//Note: Verify that comparison cannot happen if we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_EQ, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestNEQ(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic comparison operation.
	pauVM.AddInstruction(INST_PUSH, 19)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_NEQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}

	//Note: Verify basic comparison operation, but opposite.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 19)	
	pauVM.AddInstruction(INST_PUSH, 19)	
	pauVM.AddInstruction(INST_NEQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}

	//Note: Verify that comparison cannot happen if we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_NEQ, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestLS(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic less comparison.
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

	//Note: Verify basic less comparison, but the opposite.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_LS, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}

	//Note: Verify that comparison cannot happen if we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_LS, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestGR(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic greater comparison.
	pauVM.AddInstruction(INST_PUSH, 10)	
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_GR, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}

	//Note: Verify basic greater comparison, but the opposite.
	pauVM.Reset()
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 10)	
	pauVM.AddInstruction(INST_GR, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}

	//Note: Verify that comparison cannot happen if we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_GR, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestGREQ(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic greq comparison(Equality).
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

	//Note: Verify basic greq comparison(Greater).
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, -7)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_GREQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}

	//Note: Verify basic greq comparison(Less).
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 7)	
	pauVM.AddInstruction(INST_PUSH, -2)	
	pauVM.AddInstruction(INST_GREQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}

	//Note: Verify that comparison cannot happen if we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_GREQ, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestLSEQ(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic lseq comparison(Equality)
	pauVM.AddInstruction(INST_PUSH, -2)	
	pauVM.AddInstruction(INST_PUSH, -2)	
	pauVM.AddInstruction(INST_LSEQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}

	//Note: Verify basic lseq comparison(Less)
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 15)	
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_LSEQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 0 { 
		t.Errorf("Expected 0 got %d", result)
	}

	//Note: Verify basic lseq comparison(Greater)
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 3)	
	pauVM.AddInstruction(INST_PUSH, 542)	
	pauVM.AddInstruction(INST_LSEQ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 1 { 
		t.Errorf("Expected 1 got %d", result)
	}

	//Note: Verify that comparison cannot happen if we don't have at least 2 elements on the stack.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 2)
	pauVM.AddInstruction(INST_GREQ, 0)
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestDUP(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Test basic dup operation.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_DUP, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 { 
		t.Errorf("Expected 5 got %d", result)
	}

	//Note: Verify that we cannot duplicate if we don't have at least one value on the stack.
	pauVM.Reset()	

	pauVM.AddInstruction(INST_DUP, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestJMP(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic jumping.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_JMP, 6) //<-Jumps to HALT	

	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	

	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 { 
		t.Errorf("Expected 5 got %d", result)
	}

	//Note: Verify that we can't jump at an illegal place.
	pauVM.Reset()
	pauVM.AddInstruction(INST_JMP, 6) //<-Jumps at an illegal place.
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected IllegalJumpError")
	}
}

func TestJMPZ(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic jmpz.
	pauVM.AddInstruction(INST_PUSH, 123)	
	pauVM.AddInstruction(INST_PUSH, 0)	
	pauVM.AddInstruction(INST_JMPZ, 7)	

	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	

	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 123 { 
		t.Errorf("Expected 123 got %d", result)
	}

	//Note: Verify basic jmpz, but opposite.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 123)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_JMPZ, 4)	
	pauVM.AddInstruction(INST_HALT, 0) //<- SHould terminate here.

	pauVM.AddInstruction(INST_PUSH, 666)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 123 { 
		t.Errorf("Expected 123 got %d", result)
	}

	//Note: Verify that we can't jump at an illegal place.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 0)
	pauVM.AddInstruction(INST_JMPZ, 6) //<-Jumps at an illegal place.
	pauVM.AddInstruction(INST_HALT, 0)

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected IllegalJumpError")
	}

	//Note: Verify that we cannot jump if we don't have at least one value on the stack.
	pauVM.Reset()	

	pauVM.AddInstruction(INST_JMPZ, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestSWAP(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic swapping.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_PUSH, 2)	
	pauVM.AddInstruction(INST_SWAP, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]
	prevresult := pauVM.stack[sp - 1]

	if result != 5  && prevresult != 2 { 
		t.Errorf("Expected (5, 2) got (%d, %d)", result, prevresult)
	}

	//Note: Verify that we cannot swap if we don't have at least two values on the stack.
	pauVM.Reset()	

	pauVM.AddInstruction(INST_SWAP, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestSTORE(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Validate basic storing operation.
	pauVM.AddInstruction(INST_PUSH, -54)	
	pauVM.AddInstruction(INST_STORE, 1)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	result := pauVM.frames[0].locals[1]

	if result != -54 { 
		t.Errorf("Expected -54 got %d", result)
	}

	//Note: Validate storing in a random spot.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, -54)	
	pauVM.AddInstruction(INST_STORE, 8)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	result = pauVM.frames[0].locals[8]

	if result != -54 { 
		t.Errorf("Expected -54 got %d", result)
	}

	//Note: Validate that locals are different between functions.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, -54)	
	pauVM.AddInstruction(INST_STORE, 8)	
	pauVM.AddInstruction(INST_CALL, 4)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.AddInstruction(INST_PUSH, 56) //<- Function.
	pauVM.AddInstruction(INST_STORE, 2)
	pauVM.AddInstruction(INST_RETURN, 0)

	pauVM.ExecuteProgram()

	result = pauVM.frames[0].locals[8]
	result2 := pauVM.frames[1].locals[2]

	if result != -54 && result2 != 56 { 
		t.Errorf("Expected frame0: -54, frame1:  got frame0: %d, frame1: %d", result, result2)
	}

	//Note: Verify that we cannot store if we don't have at least one value on the stack.
	pauVM.Reset()	

	pauVM.AddInstruction(INST_STORE, 0)	
	pauVM.AddInstruction(INST_HALT, 0)	

	if err := pauVM.ExecuteProgram(); err == nil {
		t.Error("Expected StackUnderflowError")
	}
}

func TestLOAD(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic loading operation.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_STORE, 1)	
	pauVM.AddInstruction(INST_LOAD, 1)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 { 
		t.Errorf("Expected 5 got %d", result)
	}

	//Note: Validate loading from a random spot.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 43)	
	pauVM.AddInstruction(INST_STORE, 8)	
	pauVM.AddInstruction(INST_LOAD, 8)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 43 { 
		t.Errorf("Expected 43 got %d", result)
	}

	//Not: Validate Loading in another function.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 54)	
	pauVM.AddInstruction(INST_STORE, 8)	
	pauVM.AddInstruction(INST_CALL, 4)	
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.AddInstruction(INST_PUSH, 56) //<- Function.
	pauVM.AddInstruction(INST_STORE, 2)
	pauVM.AddInstruction(INST_LOAD, 2)
	pauVM.AddInstruction(INST_RETURN, 0)

	pauVM.ExecuteProgram()

	sp = pauVM.sp
	result = pauVM.stack[sp]

	if result != 56 { 
		t.Errorf("Expected 56 got %d", result)
	}
}

func TestCALL(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic function calling.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_CALL, 3)	
	pauVM.AddInstruction(INST_PUSH, 1)	

	pauVM.AddInstruction(INST_PUSH, 2) //<-Function
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	fp := pauVM.fp

	if fp != 1 { 
		t.Errorf("Expected 1 got %d", fp)
	}

	//Note: Verify that a function can call a function.
	pauVM.Reset()

	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_CALL, 3)	
	pauVM.AddInstruction(INST_PUSH, 1)	

	pauVM.AddInstruction(INST_PUSH, 2) //<-Function
	pauVM.AddInstruction(INST_CALL, 5)	

	pauVM.AddInstruction(INST_PUSH, 4)
	pauVM.AddInstruction(INST_HALT, 4)

	pauVM.ExecuteProgram()

	fp = pauVM.fp

	if fp != 2 { 
		t.Errorf("Expected 2 got %d", fp)
	}
}

func TestRETURN(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic return.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_CALL, 5)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	

	pauVM.AddInstruction(INST_PUSH, 2) //<- Function
	pauVM.AddInstruction(INST_RETURN, 0)
	pauVM.AddInstruction(INST_HALT, 0)	

	pauVM.ExecuteProgram()

	fp := pauVM.fp

	if fp != 0 { 
		t.Errorf("Expected 1 got %d", fp)
	}
}

func TestHALT(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Verify halting.
	pauVM.AddInstruction(INST_PUSH, 5)	
	pauVM.AddInstruction(INST_HALT, 0)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 1)	
	pauVM.AddInstruction(INST_PUSH, 2)

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 5 { 
		t.Errorf("Expected 5 got %d", result)
	}
}

func TestNONE(t *testing.T) {
	var flags VMFlags = VMFlags{
		trace: false,
		stackSize: defaultStackSize,
		callstackSize: defaultCallstackSize,
	}

	var pauVM *VM = InitVM(&flags)

	//Note: Verify basic NONE operation.
	pauVM.AddInstruction(INST_PUSH, 100)	

	pauVM.AddInstruction(INST_NONE, 1)	
	pauVM.AddInstruction(INST_NONE, 1)	
	pauVM.AddInstruction(INST_NONE, 1)	
	pauVM.AddInstruction(INST_NONE, 1)	
	pauVM.AddInstruction(INST_NONE, 1)	
	pauVM.AddInstruction(INST_NONE, 1)	
	pauVM.AddInstruction(INST_NONE, 1)	

	pauVM.AddInstruction(INST_HALT, 0)

	pauVM.ExecuteProgram()

	sp := pauVM.sp
	result := pauVM.stack[sp]

	if result != 100 { 
		t.Errorf("Expected 100 got %d", result)
	}
}
