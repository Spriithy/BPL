package vm

// Used to represent the Inner VM stack
type Stack struct {
	data []VirtValue // Stack Data
	sp   int         // Stack Pointer
}

// Simply constructs a new Emty Stack
//     s.data     will be an empty slice of VirtValues
//     s.sp       is set to -1
func NewStack() *Stack {
	return &Stack{*new([]VirtValue), -1}
}

// Peeks the top of the stack without poping it
func (s *Stack) Peek() VirtValue {
	return s.data[s.sp] // Simple look-up
}

func (s *Stack) Push(v VirtValue) {
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

func (s *Stack) Pop() VirtValue {
	v := s.data[s.sp]
	s.sp--
	return v
}

func (s *Stack) PopI() int64 {
	return s.Pop().ToInt()
}

func (s *Stack) PopR() float64 {
	return s.Pop().ToReal()
}

func (s *Stack) SP() int {
	return s.sp
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