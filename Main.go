package main

import (
	//"github.com/Spriithy/BPL/compiler"
	"github.com/Spriithy/BPL/compiler/lexer"
	"fmt"
)

func main() {

	//lx := compiler.NewLexer("main.bpl")
	//lx.Lex()

	//ps := compiler.NewParser(lx.Tokens)
	//ps.Parse()

	lxr := lexer.Lexer("main.bpl")
	lxr.Lex()

	p := lxr.List.Next
	for ; p != nil; p = p.Next {
		fmt.Println(p.String())
	}

}