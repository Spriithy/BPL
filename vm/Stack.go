package vm

import (
	"fmt"
	"log"
)

type RTStackEl struct {
	obj  RTObject
	next *RTStackEl
}

type RTStack struct {
	top  *RTStackEl
	size int
}

func InitStack() *RTStack {
	s := new(RTStack)
	s.top = nil
	s.size = 0
	return s
}

func (s *RTStack) Size() int {
	return s.size
}

func (s *RTStack) Peek() RTObject {
	return s.top.obj
}

func (s *RTStack) Pop() RTObject {
	if s.Size() > 0 {
		v := s.top.obj
		s.top = s.top.next
		s.size--
		return v
	}
	return nil
}

func (s *RTStack) Push(obj RTObject) {
	s.top = &RTStackEl{
		obj: obj,
		next: s.top,
	}
	s.size++
}

func (s *RTStack) Set(ofs int, obj RTObject) {
	if ofs < 0 || ofs > s.Size() {
		log.Fatal(fmt.Errorf("VM-Error: Accessing Stack value with invalid offset (offset = %d, size = %d)", ofs, s.Size()))
	}

	p := s.top
	for i := s.Size() - ofs - 1; i > 0; i-- {
		p = p.next
	}
	p.obj = obj
}

func (s *RTStack) Get(ofs int) RTObject {
	if ofs < 0 || ofs > s.Size() {
		log.Fatal(fmt.Errorf("VM-Error: Accessing Stack value with invalid offset (offset = %d, size = %d)", ofs, s.Size()))
	}

	p := s.top
	for i := s.Size() - ofs - 1; i > 0; i-- {
		p = p.next
	}
	return p.obj
}

func (s *RTStack) Allocate(size int) {
	for ; size > 0; size-- {
		s.Push(InitNone())
	}
}

func (s *RTStack) Free(size int) {
	for ; size > 0; size-- {
		s.Pop()
	}
}

func (s *RTStack) String() string {
	str := ""
	if s.Size() > 0 {
		if s.top.obj.Type() == None {
			str = fmt.Sprint("[None")
		} else {
			str = fmt.Sprintf("[%v", s.top.obj.RTValue())
		}
		p := s.top.next
		for p != nil {
			if p.obj.Type() == None {
				str = fmt.Sprintf("%s, None", str)
			} else {
				str = fmt.Sprintf("%s, %v", str, p.obj.RTValue())
			}
			p = p.next
		}
	} else {
		str = "Stack ["
	}

	return str + "]"
}