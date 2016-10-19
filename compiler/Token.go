package compiler

import "fmt"

type TokenType int

const (
	tEOF TokenType = iota
	tNL
	UNKNOWN

	KEYWORD
	IDENTIFIER
	NUMBER
	CHAR
	STRING

	LBRACE
	RBRACE
	LPAREN
	RPAREN
	LBRACK
	RBRACK

	DOT
	ELLIPSIS
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

var tokMEANING = []string{
	tEOF:       "~~~~~~~~~~~~~~",
	tNL:        "--------------",
	UNKNOWN:    "Unknown",

	KEYWORD:    "Keyword",
	IDENTIFIER: "Identifier",
	NUMBER:     "Number",
	CHAR:       "Char",
	STRING:     "String",

	LBRACE:     "LBrace",
	RBRACE:     "RBrace",
	LPAREN:     "LParen",
	RPAREN:     "RParen",
	LBRACK:     "LBracket",
	RBRACK:     "RBracket",

	DOT:        "Dot",
	ELLIPSIS:   "Ellipsis",
	COMMA:      "Comma",
	COLON:      "Colon",
	SEMICOL:    "Semicolon",

	ASSIGN:     "Assign",
	ARROW:      "Arrow",
	COLCOL:     "ColCol",
	OPASSIGN:   "OpAssign",
	PLUS:       "Plus",
	PLUSPLUS:   "PlusPlus",
	MINUS:      "Minus",
	MINUSMINUS: "MinusMinus",
	TIMES:      "Times",
	POW:        "Pow",
	DIV:        "Div",
	EUCL:       "Eucl",
	MOD:        "Mod",
	LSFT:       "Lshift",
	RSFT:       "Rshift",
	GT:         "Gt",
	GEQ:        "Geq",
	LT:         "Lt",
	LEQ:        "Leq",
	EQUALS:     "Equals",
	NEQUALS:    "Nequals",
	NOT:        "Not",
	AND:        "And",
	OR:         "Or",
	bNOT:       "Bnot",
	bAND:       "Band",
	bOR:        "Bor",
	bXOR:       "Bxor",
}

type Token struct {
	sym  string
	lno  int
	typ  TokenType

	Next *Token
}

func EOF(lno int) *Token {
	return &Token{"EOF", lno, tEOF, nil}
}

func NL(lno int) *Token {
	return &Token{"NewLine", lno, tNL, nil}
}

func Keyword(sym string, lno int) *Token {
	return &Token{sym, lno, KEYWORD, nil}
}

func Identifier(sym string, lno int) *Token {
	return &Token{sym, lno, IDENTIFIER, nil}
}

func Char(sym byte, lno int) *Token {
	return &Token{string(int(sym)), lno, CHAR, nil}
}

func Number(sym string, lno int) *Token {
	return &Token{sym, lno, NUMBER, nil}
}

func String(sym string, lno int) *Token {
	return &Token{sym, lno, STRING, nil}
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
	case "...": typ = ELLIPSIS
	}
	return &Token{sym, lno, typ, nil}
}

func Punctuation(sym byte, lno int) *Token {
	typ := UNKNOWN
	switch sym {
	case ',': typ = COMMA
	case ':': typ = COLON
	case ';': typ = SEMICOL
	case '(': typ = LPAREN
	case ')': typ = RPAREN
	case '[': typ = LBRACK
	case ']': typ = RBRACK
	case '{': typ = LBRACE
	case '}': typ = RBRACE
	}
	return &Token{string(sym), lno, typ, nil}
}

func (t *Token) Type() TokenType {
	return t.typ
}

func (t *Token) Symbol() string {
	return t.sym
}

func (t *Token) String() string {
	if t.typ == STRING || t.typ == IDENTIFIER {
		return fmt.Sprintf("%-3d %-16v%#v", t.lno, tokMEANING[t.typ], t.sym)
	}

	if t.typ == CHAR {
		return fmt.Sprintf("%-3d %-16v%#v", t.lno, tokMEANING[t.typ], t.sym)
	}

	return fmt.Sprintf("%-3d %-16v%s", t.lno, tokMEANING[t.typ], t.sym)
}