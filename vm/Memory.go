package vm

import (
	"fmt"
	"log"
)

type Address int

type Memory struct {
	mem map[Address]*NativeValue
	lst Address
}

func NewMemory() *Memory {
	m := new(Memory)
	m.mem = make(map[Address]*NativeValue)
	m.lst = Address(0)
	return m
}

func (m *Memory) Alloc(size int) Address {
	addr := m.lst
	m.lst = addr + Address(size)

	for i := addr; i < m.lst; i++ {
		m.mem[i] = NewVoid()
	}

	return Address(addr)
}

func (m *Memory) Free(addr Address, size int) {
	for i := addr; i < addr + Address(size); i++ {
		m.mem[i] = nil
	}

	if m.lst == addr + Address(size) {
		m.lst = addr
	}
}

func (m *Memory) Read(addr Address) *NativeValue {
	v := m.mem[addr]

	if v == nil {
		log.Fatal("Segmentation fault: 11")
	}

	return v
}

func (m *Memory) Write(addr Address, o *NativeValue) {
	m.mem[addr] = o
}

func (m *Memory) String() string {
	s := ""
	i := 1

	for addr, val := range m.mem {
		if val == nil {
			continue
		} else {
			s += fmt.Sprintf("0x%-8X\t%-16s", addr, val.String())
		}

		if i % 2 == 0 {
			s += "\n"
		}

		i++
	}

	return s
}