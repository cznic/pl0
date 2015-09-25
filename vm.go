// Copyright 2015 The PL0 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pl0

import (
	"fmt"
	"io"
	"os"

	"github.com/cznic/mathutil"
	"github.com/cznic/xc"
)

var (
	_ Instruction = (*Add)(nil)
	_ Instruction = (*Call)(nil)
	_ Instruction = (*Div)(nil)
	_ Instruction = (*Enter)(nil)
	_ Instruction = (*Halt)(nil)
	_ Instruction = (*Jmp)(nil)
	_ Instruction = (*JmpZero)(nil)
	_ Instruction = (*Leave)(nil)
	_ Instruction = (*Mul)(nil)
	_ Instruction = (*Neg)(nil)
	_ Instruction = (*PopVar)(nil)
	_ Instruction = (*PushConst)(nil)
	_ Instruction = (*PushVar)(nil)
	_ Instruction = (*Read)(nil)
	_ Instruction = (*Sub)(nil)
	_ Instruction = (*TestEQ)(nil)
	_ Instruction = (*TestGEQ)(nil)
	_ Instruction = (*TestGT)(nil)
	_ Instruction = (*TestLEQ)(nil)
	_ Instruction = (*TestLT)(nil)
	_ Instruction = (*TestMod)(nil)
	_ Instruction = (*TestOdd)(nil)
	_ Instruction = (*Write)(nil)
)

// Instruction must be implemented by all instructions of VM.
type Instruction interface {
	fmt.Stringer
	setAddr(int)
	Execute(*VM)
}

// Addr represents and address of the VM.
type Addr int

func (a *Addr) setAddr(n int) { *a = Addr(n) }

// Add is an Instruction adding the top two stack items.
type Add struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *Add) String() string {
	return fmt.Sprintf("%08d:\tadd\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *Add) Execute(m *VM) { m.push(m.pop() + m.pop()) }

// Call is an Instruction for calling subroutines.
type Call struct {
	Addr
	Target int
	Token  xc.Token
}

// String implements fmt.Stringer.
func (n *Call) String() string {
	return fmt.Sprintf("%08d:\tcall\t\t%08d\t// %s", n.Addr, n.Target, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *Call) Execute(m *VM) {
	m.push(m.IP)
	m.IP = n.Target
}

// Div is an Instruction dividing the top two stack items.
type Div struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *Div) String() string {
	return fmt.Sprintf("%08d:\tdiv\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *Div) Execute(m *VM) {
	b := m.pop()
	if b == 0 {
		panic(fmt.Errorf("division by zero"))
	}

	m.push(m.pop() / b)
}

// Enter is an Instruction setting up VM.FP and the stack frame upon entry to a
// subroutine.
type Enter struct {
	Addr
	LNL   int // Lexical nesting level.
	NVars int
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *Enter) String() string {
	return fmt.Sprintf("%08d:\tenter\t\t%v, %v\t\t// %s:%s", n.Addr, n.NVars, n.LNL, position(n.Token.Pos()), phelp(n.Token))
}

// Execute implements Instruction.
//
// Stack frame before Execute:
//
//	        +----------------+
//	        | Return address |
//	        +----------------+
//	  SP -> |                |
//
//
// Stack frame after Execute:
//
//	        +----------------+
//	        | Return address |
//	        +----------------+
//	        | Previous FP    |
//	        +----------------+
//	        | [LNL]int       | Parent Frame Pointers
//	        +----------------+
//	        | LNL            | Lexical Nesting Level
//	        +----------------+
//	  FP -> | [NVars]int     | Local variables
//	        +----------------+
//	  SP -> |                ! Evaluation stack
func (n *Enter) Execute(m *VM) {
	m.push(m.FP)
	cur := 0
	if n.LNL != 0 {
		cur = m.read(m.FP - 1)
		fp0x := m.FP - 1 - cur
		for i := 0; i < mathutil.Min(cur, n.LNL); i++ {
			m.push(m.read(fp0x))
			fp0x++
		}
	}
	if n.LNL > cur {
		m.push(m.FP)
	}
	m.push(n.LNL)
	m.FP = m.SP()
	m.Stack = append(m.Stack, make([]int, n.NVars)...)
}

// Halt is an intruction stopping the VM.
type Halt struct {
	Addr
}

// String implements fmt.Stringer.
func (n *Halt) String() string {
	return fmt.Sprintf("%08d:\thalt", n.Addr)
}

// Execute implements Instruction.
func (n *Halt) Execute(m *VM) { m.Halted = true }

// Jmp is an Instruction uncoditionally changing VM.IP.
type Jmp struct {
	Addr
	Target int
}

// String implements fmt.Stringer.
func (n *Jmp) String() string { return fmt.Sprintf("%08d:\tjmp\t\t%08d", n.Addr, n.Target) }

// Execute implements Instruction.
func (n *Jmp) Execute(m *VM) { m.IP = n.Target }

// JmpZero is an Instruction coditionally changing VM.IP when TOS == 0.
type JmpZero struct {
	Addr
	Target int
}

// String implements fmt.Stringer.
func (n *JmpZero) String() string { return fmt.Sprintf("%08d:\tjz\t\t%08d", n.Addr, n.Target) }

// Execute implements Instruction.
func (n *JmpZero) Execute(m *VM) {
	if m.pop() == 0 {
		m.IP = n.Target
	}
}

// Leave is an Instruction restoring the stack frame and IP prior to a
// subroutine call.
type Leave struct {
	Addr
}

// String implements fmt.Stringer.
func (n *Leave) String() string { return fmt.Sprintf("%08d:\tleave", n.Addr) }

// Execute implements Instruction.
func (n *Leave) Execute(m *VM) {
	m.Stack = m.Stack[:m.FP]
	lnl := m.pop()
	m.Stack = m.Stack[:m.SP()-lnl]
	m.FP = m.pop()
	m.IP = m.pop()
}

// Mul is an Instruction multiplying the top two stack items.
type Mul struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *Mul) String() string {
	return fmt.Sprintf("%08d:\tmul\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *Mul) Execute(m *VM) { m.push(m.pop() * m.pop()) }

// Neg is an Instruction multiplying TOS by -1.
type Neg struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *Neg) String() string {
	return fmt.Sprintf("%08d:\tsub\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *Neg) Execute(m *VM) { m.push(-m.pop()) }

// PopVar is an Instruction popping a variable on stack.
type PopVar struct {
	Addr
	Index  int
	Frames int
	Token  xc.Token
}

// String implements fmt.Stringer.
func (n *PopVar) String() string {
	switch n.Frames {
	case 0:
		return fmt.Sprintf("%08d:\tpopVar\t\t%v\t\t// %s:%s", n.Addr, n.Index, position(n.Token.Pos()), phelp(n.Token))
	default:
		return fmt.Sprintf("%08d:\tpopVar\t\t%v, %v\t\t// %s:%s", n.Addr, n.Frames, n.Index, position(n.Token.Pos()), phelp(n.Token))
	}
}

// Execute implements Instruction.
func (n *PopVar) Execute(m *VM) {
	fp := m.fp(n.Frames)
	m.write(fp+n.Index, m.pop())
}

// PushConst is an Instruction pushing a constant on stack.
type PushConst struct {
	Addr
	Value int
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *PushConst) String() string {
	return fmt.Sprintf("%08d:\tpushConst\t%v\t\t// %s", n.Addr, n.Value, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *PushConst) Execute(m *VM) { m.push(n.Value) }

// PushVar is an Instruction pushing a variable on stack.
type PushVar struct {
	Addr
	Index  int
	Frames int
	Token  xc.Token
}

// String implements fmt.Stringer.
func (n *PushVar) String() string {
	switch n.Frames {
	case 0:
		return fmt.Sprintf("%08d:\tpushVar\t\t%v\t\t// %s:%s", n.Addr, n.Index, position(n.Token.Pos()), phelp(n.Token))
	default:
		return fmt.Sprintf("%08d:\tpushVar\t\t%v, %v\t\t// %s:%s", n.Addr, n.Frames, n.Index, position(n.Token.Pos()), phelp(n.Token))
	}
}

// Execute implements Instruction.
func (n *PushVar) Execute(m *VM) {
	fp := m.fp(n.Frames)
	m.push(m.read(fp + n.Index))
}

// Read is an Instruction setting a variable from user input.
type Read struct {
	Addr
	Index  int
	Frames int
	Token  xc.Token
}

// String implements fmt.Stringer.
func (n *Read) String() string {
	switch n.Frames {
	case 0:
		return fmt.Sprintf("%08d:\tread\t\t%v\t\t// %s:%s", n.Addr, n.Index, position(n.Token.Pos()), phelp(n.Token))
	default:
		return fmt.Sprintf("%08d:\tread\t\t%v, %v\t\t// %s:%s", n.Addr, n.Frames, n.Index, position(n.Token.Pos()), phelp(n.Token))
	}
}

// Execute implements Instruction.
func (n *Read) Execute(m *VM) {
	fmt.Fprintf(m.Stdout, "?")
	var val int
	fmt.Fscanln(m.Stdin, &val)
	fp := m.fp(n.Frames)
	m.write(fp+n.Index, val)
}

// Sub is an Instruction producing the difference of the top two stack items.
type Sub struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *Sub) String() string {
	return fmt.Sprintf("%08d:\tsub\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *Sub) Execute(m *VM) {
	b := m.pop()
	m.push(m.pop() - b)
}

// TestEQ is an Instruction testing equality of the top two stack items.
type TestEQ struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *TestEQ) String() string {
	return fmt.Sprintf("%08d:\ttest =\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *TestEQ) Execute(m *VM) { m.push(b2i[m.pop() == m.pop()]) }

// TestGEQ is an Instruction testing >= of the top two stack items.
type TestGEQ struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *TestGEQ) String() string {
	return fmt.Sprintf("%08d:\ttest >=\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *TestGEQ) Execute(m *VM) {
	b := m.pop()
	m.push(b2i[m.pop() >= b])
}

// TestGT is an Instruction testing > of the top two stack items.
type TestGT struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *TestGT) String() string {
	return fmt.Sprintf("%08d:\ttest >\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *TestGT) Execute(m *VM) {
	b := m.pop()
	m.push(b2i[m.pop() > b])
}

// TestLEQ is an Instruction testing <= of the top two stack items.
type TestLEQ struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *TestLEQ) String() string {
	return fmt.Sprintf("%08d:\ttest <=\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *TestLEQ) Execute(m *VM) {
	b := m.pop()
	m.push(b2i[m.pop() <= b])
}

// TestLT is an Instruction testing < of the top two stack items.
type TestLT struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *TestLT) String() string {
	return fmt.Sprintf("%08d:\ttest <\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *TestLT) Execute(m *VM) {
	b := m.pop()
	m.push(b2i[m.pop() < b])
}

// TestMod is an Instruction testing if the top two stack items have non zero reminder on division.
type TestMod struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *TestMod) String() string {
	return fmt.Sprintf("%08d:\ttest #\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *TestMod) Execute(m *VM) {
	b := m.pop()
	if b == 0 {
		panic(fmt.Errorf("division by zero"))
	}
	m.push(b2i[m.pop()%b != 0])
}

// TestOdd is an Instruction testing if TOS&1 != 0.
type TestOdd struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *TestOdd) String() string {
	return fmt.Sprintf("%08d:\ttest odd\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *TestOdd) Execute(m *VM) { m.push(b2i[m.pop()&1 != 0]) }

// Write is an Instruction writing TOS to VM.Stdout.
type Write struct {
	Addr
	Token xc.Token
}

// String implements fmt.Stringer.
func (n *Write) String() string {
	return fmt.Sprintf("%08d:\twrite\t\t\t\t// %s", n.Addr, position(n.Token.Pos()))
}

// Execute implements Instruction.
func (n *Write) Execute(m *VM) {
	if _, err := fmt.Fprintln(m.Stdout, m.pop()); err != nil {
		panic(err)
	}
}

// Run executes n.Code. opts optionally ammend the VM used to execute the program.
func (n *Program) Run(opts ...RunOption) error {
	vm := NewVM(os.Stdin, os.Stdout, n.Code)
	for _, opt := range opts {
		if err := opt(vm); err != nil {
			return err
		}
	}
	return vm.Run()
}

// List list n.Code to w.
func (n *Program) List(w io.Writer) error {
	for i, v := range n.Code {
		if _, err := fmt.Fprintln(w, v); err != nil {
			return err
		}

		switch v.(type) {
		case *Jmp, *JmpZero, *Leave, *Halt:
			if i != len(n.Code)-1 {
				if _, err := fmt.Fprintln(w); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// RunOption is an option of (*Program).Run.
type RunOption func(*VM) error

// Trace turns on tracing of instructions executed.
func Trace() RunOption {
	return func(v *VM) error {
		v.Trace = true
		return nil
	}
}

// TraceStack turns on tracing of instructions executed and stack state.
func TraceStack() RunOption {
	return func(v *VM) error {
		v.TraceStack = true
		return nil
	}
}

// VM is PL/0 virtual machine.
type VM struct {
	Code       []Instruction
	FP         int
	Halted     bool
	IP         int
	Stack      []int
	Stdin      io.Reader
	Stdout     io.Writer
	T          int
	Trace      bool
	TraceStack bool
}

// NewVM returns a newly created VM.
func NewVM(stdin io.Reader, stdout io.Writer, code []Instruction) *VM {
	return &VM{
		Code:   code,
		Stdin:  stdin,
		Stdout: stdout,
	}
}

// Run executes m.Code.
func (m *VM) Run() (err error) {
	for !m.Halted {
		if err := m.Step(); err != nil {
			return err
		}
	}
	return nil
}

// Step executes on instruction of m.
func (m *VM) Step() (err error) {
	if m.IP >= len(m.Code) {
		return fmt.Errorf("%08d: segfault on instruction fetch", m.IP)
	}

	defer func() {
		if e := recover(); e != nil && err == nil {
			err = fmt.Errorf("%v", e)
		}
	}()

	inst := m.Code[m.IP]
	m.IP++
	m.T++
	switch {
	case m.TraceStack:
		fmt.Fprintf(m.Stdout, "T %04d,  FP %04d, SP %04d, %v:\t%v\n", m.T, m.FP, m.SP(), m.Stack, inst)
	case m.Trace:
		fmt.Fprintf(m.Stdout, "T %04d,  FP %04d, SP %04d:\t%v\n", m.T, m.FP, m.SP(), inst)
	}
	inst.Execute(m)
	return nil
}

// SP returns the stack pointer of m.
func (m *VM) SP() int { return len(m.Stack) }

func (m *VM) push(n int) { m.Stack = append(m.Stack, n) }

func (m *VM) pop() int {
	sp := m.SP() - 1
	r := m.read(sp)
	m.Stack = m.Stack[:sp]
	return r
}

func (m VM) read(off int) int {
	if off >= m.SP() {
		panic(fmt.Errorf("%08d: segfault on read, index %v", m.IP, off))
	}

	return m.Stack[off]
}

func (m VM) write(off, val int) {
	if off >= m.SP() {
		panic(fmt.Errorf("%08d: segfault on write, index %v", m.IP, off))
	}

	m.Stack[off] = val
}

func (m *VM) fp(frames int) int {
	if frames == 0 {
		return m.FP
	}

	return m.read(m.FP - 2 - (frames - 1))
}
