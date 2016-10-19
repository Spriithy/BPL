package main

import (
	"./compiler"
)

func main() {
	/*	s := vm.NewStack()

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

		p := new([]vm.VMInstruction)
		(*p)[0] = vm.Push{vm.NewInt(10)}
		(*p)[1] = vm.Dec{}
		(*p)[2] = vm.Halt{}

		m := vm.NewMachine(p)
		m.Start()

		*/
	lx := compiler.NewLexer("main.bpl")
	lx.Lex()

	ps := compiler.NewParser(lx.Tokens)
	ps.Parse()
}