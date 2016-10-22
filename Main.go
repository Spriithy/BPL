package main

import (
	"github.com/Spriithy/BPL/compiler/lexer"
	"github.com/Spriithy/BPL/compiler/parser"
	"github.com/Spriithy/BPL/vm"
)

func main() {

	lxr := lexer.Lexer("main.bpl")
	lxr.Lex()

	p := lxr.List.PeekHead()
	for ; p != nil; p = p.Next {
		println(p.String())
	}

	println("-----------------------------")

	pr := parser.Parser(lxr.Path, lxr.Source(), lxr.List)
	pr.Parse()

	println("-----------------------------")

	v := vm.VirtualMachine([]vm.Bytecode{
		vm.ICONST_N, 'A', vm.ICONST_N, 'B',
		vm.IEQ, vm.PRINTLN_VAL,
		vm.HALT,
	})

	println("    IP    |  INSTRUCTION  | ARGS")
	v.Start()

}