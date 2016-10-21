package token

import "fmt"

type Pos struct {
	Start, End int // Both inclusive
}

func (p *Pos) String() string {
	return fmt.Sprintf("(%d,%d)", p.Start, p.End)
}

type Token struct {
	Kind, Sym string
	Pos, Line Pos
	Lno       int
	Next      *Token
}

func (t *Token) String() string {
	vfmt := "%-s"
	sym := t.Sym
	if t.Kind == "String" {
		vfmt = "%-#v"
	} else if t.Kind == "Character" || t.Kind == "Identifier" {
		vfmt = "'%-s'"
		sym = fmt.Sprintf("%#v", t.Sym)
		sym = sym[1:len(sym) - 1]
	}
	return fmt.Sprintf("[%4d]\t%-16s " + vfmt, t.Lno, t.Kind, sym)
}

func (t *Token) ShortString() string {
	vfmt := "%-s"
	sym := t.Sym
	if t.Kind == "String" {
		vfmt = "%-#v"
	} else if t.Kind == "Character" || t.Kind == "Identifier" {
		vfmt = "'%-s'"
		sym = fmt.Sprintf("%#v", t.Sym)
		sym = sym[1:len(sym) - 1]
	}
	return fmt.Sprintf("%-16s " + vfmt, t.Kind, sym)
}