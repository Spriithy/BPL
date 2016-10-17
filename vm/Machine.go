package vm

type Machine struct {
	prg     []*VMInstruction
	ip      int

	stk     Stack
	mem     Memory
	running bool
}

func NewMachine(prg []VMInstruction) *Machine {
	m := new(Machine)
	m.stk = *NewStack()
	m.mem = *NewMemory()
	m.prg = prg
	m.ip = 0
	m.running = false
	return m
}

func (m *Machine) Start() {
	m.running = true
	for ; m.running; {
		(*m.prg[m.ip]).Exec(*m)
		m.ip++
	}
}

func (m *Machine) Halt() {
	m.running = false
}

func (m *Machine) String() string {
	return m.stk.String()
}