package main

import (
	"./vm"
	"fmt"
)

func main() {
	s := vm.NewStack()

	for i := 0; i < 20; i++ {
		s.Push(vm.NewInt(i * (i + 1)))
		fmt.Println(s.Pop().String())
	}

	s.Push(vm.NewString("Foo"))
	fmt.Println(s)

	mem := vm.NewMemory()
	mem.Alloc(19)
	mem.Write(100, vm.NewByte(10))

	fmt.Println(mem.String())
	fmt.Println(mem.Read(1))
}