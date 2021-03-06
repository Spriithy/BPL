package main

import (
	"github.com/Spriithy/BPL/compiler/lexer"
	"github.com/Spriithy/BPL/compiler/parser"
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
}