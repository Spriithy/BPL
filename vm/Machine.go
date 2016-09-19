package vm

import "fmt"

type Machine struct {
	stack RTStack
	prog  []int

	jp    int  // Jump pointer used to set IP
	rp    int  // Return Instruction Pointer
	sp    int  // Stack Pointer
	fp    int  // Frame Pointer
	ip    int  // Instruction Pointer

	state bool // Whether the machine is active or not
}

func InitMachine(prog []int, args ... RTObject) *Machine {
	m := new(Machine)
	m.stack = *InitStack()
	m.prog = prog
	m.rp = -1
	m.sp = -1
	m.ip = 0
	m.fp = 1 + len(args)

	m.state = false

	i := len(args) - 1
	for ; i >= 0; i-- {
		m.StackPush(args[i])
	}
	m.StackPush(NativeInt(int64(len(args))))
	m.StackPush(NativeInt(-1))

	return m
}

func (m *Machine) StackPush(obj RTObject) {
	m.sp++
	m.stack.Push(obj)
}

func (m *Machine) StackPop() RTObject {
	m.sp--
	return m.stack.Pop()
}

func (m *Machine) GetArg(n int) RTObject {
	return m.stack.Get(m.fp - n - 1)
}

func (m *Machine) GetLocal(n int) RTObject {
	return m.stack.Get(m.fp + n + 1)
}

func (m *Machine) SetLocal(n int) {
	m.stack.Set(m.fp + n + 1, m.StackPop())
}

func (m *Machine) Start() {
	m.state = true
	ops := SetupInstructionSet()

	for ; m.state; {
		if m.ip < len(m.prog) - 1 {
			fmt.Printf("%d:\tCurr = %d, Next = %d\n", m.ip, m.prog[m.ip], m.prog[m.ip + 1])
		}
		ops[m.prog[m.ip]](m)
		fmt.Println(&m.stack)
		m.ip++
	}
}

func (m *Machine) Halt() {
	fmt.Println("VM Halted!")
	m.state = false
}

func (m *Machine) Print() {
	fmt.Println("::======[ VIRTUAL MACHINE ]=================================================::")
	fmt.Printf("IP%16d\n", m.ip)
	fmt.Printf("SP%16d\n", m.sp)
	fmt.Printf("FP%16d\n", m.fp)
	fmt.Printf("TOS%15v (%s)\n", m.stack.Peek().RTValue(), m.stack.Peek().RTTypeStr())
	fmt.Printf("%s\n", m.stack.String())
	fmt.Println("::==========================================================================::")
}