package vm

import (
	"fmt"
	"log"
	"reflect"
)

type Instruction func(*Machine)

const (
	NOP int = iota // No op

	JMP        // Jump
	JZ        // Jump if zero
	JNZ        // Jump if not zero

	ALLOC        // Alloc TOS space on local frame

	ARG        // Get nth Argument passed ARG 0 = nargs

	LGET        // Local get
	LSET        // Local set
	GGET        // Global get
	GSET        // Global set

	PUSH        // Push next value on stack
	POP        // Pop off of stack

	DUP        // Duplicate TOS
	SWAP        // Swap TOS and TOS-1

	CALL        // Call function with ID
	CLEAR        // Clear stack of current frame
	RET        // Return from current function

	HLT        // Halts the machine
	HZ        // Halts if zero
	HNZ        // Halts if not zero

	INC        // Increments TOS by 1
	DEC        // Decrements TOS by 1

)

func SetupInstructionSet() map[int]Instruction {
	var set map[int]Instruction = make(map[int]Instruction)

	set[NOP] = func(*Machine) {}

	set[JMP] = func(m *Machine) {
		m.ip = m.prog[m.ip + 1]
		fmt.Println("Jumping at -> ", m.prog[m.ip + 1])
	}

	set[JZ] = func(m *Machine) {
		if m.StackPop().RTValue() == 0 {
			set[JMP](m)
		}
	}

	set[JNZ] = func(m *Machine) {
		if m.StackPop().RTValue() != 0 {
			set[JMP](m)
		}
	}

	set[ALLOC] = func(m *Machine) {
		p := m.prog[m.ip + 1]
		m.stack.Allocate(p)
		m.ip++
	}

	set[ARG] = func(m *Machine) {
		p := m.prog[m.ip + 1]
		m.StackPush(m.GetArg(p))
		m.ip++
	}

	set[LGET] = func(m *Machine) {
		p := m.prog[m.ip + 1]
		v := m.GetLocal(p)
		m.StackPush(v)
		m.ip++
	}

	set[LSET] = func(m *Machine) {
		p := m.prog[m.ip + 1]
		m.SetLocal(p)
		m.ip++
	}

	set[GGET] = func(m *Machine) {
		// TODO
	}

	set[GSET] = func(m *Machine) {
		// TODO
	}

	set[DUP] = func(m *Machine) {
		v := m.StackPop()
		m.StackPush(v)
		m.StackPush(v)
	}

	set[SWAP] = func(m *Machine) {
		v1 := m.StackPop()
		v2 := m.StackPop()
		m.StackPush(v1)
		m.StackPush(v2)
	}

	set[CALL] = func(m *Machine) {

	}

	set[CLEAR] = func(m *Machine) {
		for i := 0; i < m.stack.Size() - m.fp + 1; i++ {
			m.StackPop()
		}

		fp := int(m.StackPop().RTValue().(int64))

		if fp == -1 {
			set[HLT](m)
			return
		}

		m.fp = fp
	}

	set[RET] = func(m *Machine) {
		set[CLEAR](m)
	}

	set[HLT] = func(m *Machine) {
		m.Halt()
	}

	set[HZ] = func(m *Machine) {
		if m.StackPop().RTValue() == 0 {
			set[HLT](m)
		}
	}

	set[HNZ] = func(m *Machine) {
		if m.StackPop().RTValue() != 0 {
			set[HLT](m)
		}
	}

	set[PUSH] = func(m *Machine) {
		m.StackPush(InitNative(Int, m.prog[m.ip + 1]))
		m.ip++
	}

	set[POP] = func(m *Machine) {
		m.StackPop()
	}

	set[INC] = func(m *Machine) {
		p := m.StackPop()

		if p == nil {
			log.Fatal(fmt.Errorf("VM-Error:%d: Null Pointer Error on DEC operation", m.ip))
		}

		t, v := p.Type(), p.RTValue()
		switch t {
		case Int:
			v := reflect.ValueOf(v).Int()
			m.StackPush(NativeInt(v + 1))
			return
		case UInt:
			v := reflect.ValueOf(v).Uint()
			m.StackPush(NativeUInt(v + 1))
			return
		default:
			log.Fatal(fmt.Errorf("VM-Error:%d: INC operation on invalid type (%s)", m.ip, p.RTTypeStr()))
			return
		}
	}

	set[DEC] = func(m *Machine) {
		p := m.StackPop()

		if p == nil {
			log.Fatal(fmt.Errorf("VM-Error:%d: Null Pointer Error on DEC operation", m.ip))
		}

		t, v := p.Type(), p.RTValue()
		switch t {
		case Int:
			v := reflect.ValueOf(v).Int()
			m.StackPush(NativeInt(v - 1))
			return
		case UInt:
			v := reflect.ValueOf(v).Uint()
			m.StackPush(NativeUInt(v - 1))
			return
		case String:
			log.Fatal(fmt.Errorf("VM-Error:%d: INC operation on invalid type (%s)", m.ip, p.RTTypeStr()))
			return
		}
	}

	return set
}