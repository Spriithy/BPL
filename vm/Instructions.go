package vm

type VMInstruction interface {
	Exec(Machine)
}

type Nop struct {}
func (i Nop) Exec(m Machine) {}

type Halt struct {}
func (i Halt) Exec(m Machine) {
	m.Halt()
}

type Push struct { El *NativeValue }
func (i Push) Exec(m Machine) {
	m.stk.Push(i.El)
}

type Pop struct {}
func (i Pop) Exec(m Machine) {
	m.stk.Pop()
}

type Jump struct { to NativeValue }
func (i Jump) Exec(m Machine) {
	m.ip = i.to.int
}

type Jumpz struct { to NativeValue }
func (i Jumpz) Exec(m Machine) {
	if m.stk.Pop().IsZero() {
		m.ip = i.to.int
	}
}

type Jumpnz struct { to NativeValue }
func (i Jumpnz) Exec(m Machine) {
	if !m.stk.Pop().IsZero() {
		m.ip = i.to.int
	}
}

type Inc struct {}
func (i Inc) Exec(m Machine) {
	o := m.stk.Pop()
	switch o.TypeOf() {
	case ByteType:
		m.stk.Push(NewByte(o.Byte() + 1))
		return
	case IntType:
		m.stk.Push(NewInt(o.Int() + 1))
		return
	case UIntType:
		m.stk.Push(NewUInt(o.UInt() + 1))
		return
	case LongType:
		m.stk.Push(NewLong(o.Long() + 1))
		return
	case ULongType:
		m.stk.Push(NewULong(o.ULong() + 1))
		return
	case FloatType:
		m.stk.Push(NewFloat(o.Float() + 1.))
		return
	case DoubleType:
		m.stk.Push(NewDouble(o.Double() + 1.))
		return
	}
}

type Dec struct {}
func (i Dec) Exec(m Machine) {
	o := m.stk.Pop()
	switch o.TypeOf() {
	case ByteType:
		v := o.Byte()
		if v > 0 {
			m.stk.Push(NewByte(v - 1))
			return
		}
		m.stk.Push(o)
		return
	case IntType:
		m.stk.Push(NewInt(o.Int() - 1))
		return
	case UIntType:
		v := o.UInt()
		if v > 0 {
			m.stk.Push(NewUInt(v - 1))
			return
		}
		m.stk.Push(o)
		return
	case LongType:
		m.stk.Push(NewLong(o.Long() - 1))
		return
	case ULongType:
		v := o.ULong()
		if v > 0 {
			m.stk.Push(NewULong(v - 1))
			return
		}
		m.stk.Push(o)
		return
	case FloatType:
		m.stk.Push(NewFloat(o.Float() - 1.))
		return
	case DoubleType:
		m.stk.Push(NewDouble(o.Double() - 1.))
		return
	}
}
