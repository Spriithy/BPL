package vm

import (
	"os"
)

// Used to represent the Inner VM stack
type Stack struct {
	data []VirtualValue // Stack Data
	sp   int            // Stack Pointer
}

// Simply constructs a new Emty Stack
//     s.data     will be an empty slice of VirtValues
//     s.sp       is set to -1
func NewStack() *Stack {
	return &Stack{*new([]VirtualValue), -1}
}

// Peeks the top of the stack without poping it
func (s *Stack) Peek() VirtualValue {
	return s.data[s.sp] // Simple look-up
}

func (s *Stack) Push(v VirtualValue) {
	s.sp++
	if len(s.data) == s.sp {
		(*s).data = append((*s).data, v)
		return
	}
	s.data[s.sp] = v
}

func (s *Stack) PushI(i int64) {
	s.Push(VirtualInt(i))
}

func (s *Stack) PushR(r float64) {
	s.Push(VirtualReal(r))
}

func (s *Stack) Pop() VirtualValue {
	if s.Empty() {
		println("Cannot pop from empty stack!", s.String())
		os.Exit(1)
	}

	v := s.data[s.sp]
	s.sp--
	return v
}

func (s *Stack) PopI() int64 {
	v1 := s.Pop()
	if v1.Type() == VIRTUAL_NULL {
		println("NullPointer to Integer conversion!")
		os.Exit(1)
	}
	return v1.ToInt()
}

func (s *Stack) PopR() float64 {
	v1 := s.Pop()
	if v1.Type() == VIRTUAL_NULL {
		println("NullPointer to Real conversion!")
		os.Exit(1)
	}
	return v1.ToReal()
}

func (s *Stack) Empty() bool {
	return s.sp < 0
}

func (s *Stack) String() string {
	if s.Empty() {
		return "[]"
	}
	str := ""
	for i := s.sp; i >= 0; i-- {
		str += " " + s.data[i].String()
	}
	return "[" + str[1:] + "]"
}