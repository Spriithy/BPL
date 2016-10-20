package main

import (
	"github.com/Spriithy/BPL/compiler/lexer"
	"github.com/Spriithy/BPL/compiler/parser"
)

func main() {

	//lx := compiler.NewLexer("main.bpl")
	//lx.Lex()

	//ps := compiler.NewParser(lx.Tokens)
	//ps.Parse()

	lxr := lexer.Lexer("main.bpl")
	lxr.Lex()

	p := lxr.List.PeekHead()
	for ; p.Next != nil; p = p.Next {
		println(p.String())
	}

	println("-----------------------------")

	pr := parser.Parser(lxr.Path, lxr.Source(), lxr.List)
	pr.Parse()

}