package compiler

import "fmt"

type Parser struct {
	tokens, tok *Token
	at, lno     int
}

func NewParser(tokens *Token) *Parser {
	p := new(Parser)
	p.at, p.lno = 0, 0
	p.tokens = tokens
	p.tok = tokens.Next
	return p
}

func (p *Parser) Parse() {
	L := p.tokens
	for p := L.Next; p != nil; p = p.Next {
		fmt.Println(p.String())
	}
}