package compiler

import "fmt"

type Type int

const (
	tEOF Type = iota
	UNKNOWN

	KEYWORD
	IDENTIFIER
	NUMBER
	STRING
	BOOL
	NULL

	LBRACE
	RBRACE
	LPAREN
	RPAREN
	LBRACK
	RBRACK

	DOT
	DDDOT
	COMMA
	COLON
	SEMICOL

	ASSIGN
	ARROW
	COLCOL
	OPASSIGN
	PLUS
	PLUSPLUS
	MINUS
	MINUSMINUS
	TIMES
	POW
	DIV
	EUCL
	MOD
	LSFT
	RSFT
	GT
	GEQ
	LT
	LEQ
	EQUALS
	NEQUALS
	NOT
	AND
	OR
	bNOT
	bAND
	bOR
	bXOR
)

type Token struct {
	sym  string
	lno  int
	typ  Type

	Next *Token
}

func EOF(lno int) *Token {
	return &Token{"EOF", lno, tEOF, nil}
}

func Null(lno int) *Token {
	return &Token{"null", lno, NULL, nil}
}

func Keyword(sym string, lno int) *Token {
	return &Token{sym, lno, KEYWORD, nil}
}

func Identifier(sym string, lno int) *Token {
	return &Token{sym, lno, IDENTIFIER, nil}
}

func Number(sym string, lno int) *Token {
	return &Token{sym, lno, NUMBER, nil}
}

func String(sym string, lno int) *Token {
	return &Token{sym, lno, STRING, nil}
}

func Bool(sym string, lno int) *Token {
	return &Token{sym, lno, BOOL, nil}
}

func Operator(sym string, lno int) *Token {
	typ := UNKNOWN
	switch sym {
	case "=": typ = ASSIGN
	case "+": typ = PLUS
	case "-": typ = MINUS
	case "*": typ = TIMES
	case "/": typ = DIV
	case "%": typ = MOD
	case "++": typ = PLUSPLUS
	case "--": typ = MINUSMINUS
	case "**": typ = POW
	case "//": typ = EUCL
	case "<<": typ = LSFT
	case ">>": typ = RSFT

	case "==": typ = EQUALS
	case "!=": typ = NEQUALS
	case "&&": typ = AND
	case "||": typ = OR
	case ">=": typ = GEQ
	case "<=": typ = LEQ
	case "!": typ = NOT
	case ">": typ = GT
	case "<": typ = LT

	case "+=": typ = OPASSIGN
	case "-=": typ = OPASSIGN
	case "/=": typ = OPASSIGN
	case "%=": typ = OPASSIGN
	case "**=": typ = OPASSIGN
	case "//=": typ = OPASSIGN

	case "&=": typ = OPASSIGN
	case "|=": typ = OPASSIGN
	case "^=": typ = OPASSIGN
	case "<<=": typ = OPASSIGN
	case ">>=": typ = OPASSIGN

	case "~": typ = bNOT
	case "&": typ = bAND
	case "|": typ = bOR
	case "^": typ = bXOR

	case ".": typ = DOT
	case "->": typ = ARROW
	case "::": typ = COLCOL
	case "...": typ = DDDOT
	}
	return &Token{sym, lno, typ, nil}
}

func Punctuation(sym string, lno int) *Token {
	typ := UNKNOWN
	switch sym {
	case ",": typ = COMMA
	case ":": typ = COLON
	case ";": typ = SEMICOL
	case "(": typ = LPAREN
	case ")": typ = RPAREN
	case "[": typ = LBRACE
	case "]": typ = RBRACE
	case "{": typ = LBRACK
	case "}": typ = RBRACK
	}
	return &Token{sym, lno, typ, nil}
}

func (t *Token) String() string {
	return fmt.Sprintf("([%d]: %v, %#v)", t.lno, t.typ, t.sym)
}