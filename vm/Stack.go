package vm

import "fmt"

type Stack []*NativeValue

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(n *NativeValue) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (n *NativeValue) {
	x := s.Sp() - 1
	n = (*s)[x]
	*s = (*s)[:x]
	return
}

func (s *Stack) Sp() int {
	return len(*s)
}

func (s *Stack) String() string {
	str := "["

	for i := s.Sp() - 1; i >= 0; i-- {
		str += fmt.Sprintf("%v", (*s)[i].String())
		if i != 0 {
			str += ", "
		}
	}

	return str + "]"
}