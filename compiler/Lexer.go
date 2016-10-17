package compiler

import (
	"io/ioutil"
	"fmt"
)

var lxOPERATORS = []string{
	"=", "+", "-", "*", "/", "%", "~", "|", "^", "&", "<", ">", "!", "++",
	"--", "**", "//", "==", "<<", ">>", "!=", "+=", "-=", "*=", "/=", "**=",
	"//=", "<<=", ">>=", "&=", "|=", "^=", "~=", "||", "&&", "...", ".",
	"::", "->" }

var lxKEYWORDS = []string{
	"var", "const", "struct", "enum", "public", "private", "static", "new",

	"if", "else", "for", "while", "unless", "return", "switch", "match",
	"true", "false", "null", "final", "this", "break", "skip",
	"int", "long", "float", "double", "complex", "byte", "string",
	"uint", "ulong", "bool" }

type Lexer struct {
	Path, input  string
	at, lno      int
	cur, nxt     byte
	Tokens, last *Token
}

func NewLexer(path string) *Lexer {
	l := new(Lexer)
	l.Path = path
	f, e := ioutil.ReadFile(path)
	if e != nil {}
	l.input = string(f) + "\n"
	l.Tokens = EOF(-1)
	l.last = l.Tokens
	return l
}

func (l *Lexer) init() {
	l.cur = ' '
	l.nxt = ' '
	l.at = 0
	l.lno = 1
}

func (l *Lexer) Lex() {
	l.init()
	for ; l.next(); {
		l.process()
	}
	l.last.Next = EOF(l.lno)
}

func (l *Lexer) process() {
	fmt.Printf("%-5dc:%c\tn:%c\n", l.at, l.cur, l.nxt)
	c, n := string(l.cur), string(l.nxt)
	if c + n == "//" {
		l.comment(0)
	} else if c + n == "/*" {
		l.comment(1)
	} else if c == "\n" || c == "\t" || c == " " || c == "\r" {
		return
	} else if c == "\"" {
		l.dostring()
	} else if isPunctuation(l.cur) {
		l.punctuation()
	} else if isIdStart(l.cur) {
		l.identifier()
	} else if isNumStart(l.cur) {
		l.number()
	} else if isOperator(c) {
		l.operator()
	} else {
		/* TODO ERROR */
	}
}

func (l *Lexer) punctuation() {
	l.last.Next = Punctuation(string(l.cur), l.lno)
	l.last = l.last.Next
}

func (l *Lexer) operator() {
	op1 := string(l.cur)
	op2 := op1 + string(l.nxt)
	l.next()
	op3 := op2 + string(l.nxt)
	if isOperator(op3) {
		l.next()
		l.last.Next = Operator(op3, l.lno)
		l.last = l.last.Next
		return
	}

	if isOperator(op2) {
		l.last.Next = Operator(op2, l.lno)
		l.last = l.last.Next
		return
	}

	if isOperator(op1) {
		l.last.Next = Operator(op1, l.lno)
		l.last = l.last.Next
		return
	}
}

func (l *Lexer) dostring() {
	escape := false
	str := ""
	l.next()
	for ; l.cur != '"'; {
		if l.cur == '\\' {
			escape = ! escape
		}

		if escape {
			switch l.nxt {
			case 'n': str += "\n"; break
			case 't': str += "\t"; break
			case 'r': str += "\r"; break
			case '"': str += "\""; break
			case '\\': str += "\\"; break
			default:
				/* TODO ERROR */
				return
			}
			l.next()
			escape = false
		} else {
			str += string(l.cur)
		}
		l.next()
	}
	l.last.Next = String(str, l.lno)
	l.last = l.last.Next
}

func (l *Lexer) identifier() {
	id := string(l.cur)
	for ; isIdPart(l.nxt); {
		id += string(l.nxt)
		l.next()
	}

	if isKeyword(id) {
		l.last.Next = Keyword(id, l.lno)
	} else {
		l.last.Next = Identifier(id, l.lno)
	}
	l.last = l.last.Next
}

func (l *Lexer) comment(mode int) {
	const INLINE int = 0
	const MULTILINE int = 1
	switch (mode) {
	case INLINE:
		for ; l.cur != '\n'; {
			l.next();
		}
		break;
	case MULTILINE:
		for ; string(l.cur) + string(l.nxt) != "*/"; {
			l.next();
		}
		l.next()
		break;
	}
}

func (l *Lexer) number() {
	decimal := false
	num := string(l.cur)
	for ; isNumPart(l.nxt); {
		if l.nxt == '.' && !decimal {
			decimal = true
		} else if l.nxt == '.' {
			/* TODO ERROR */
		}
		num += string(l.nxt)
		l.next();
	}
	l.last.Next, l.last = Number(num, l.lno), l.last.Next
}

func isPunctuation(b byte) bool {
	return b == '.' || b == ',' || b == ':' || b == ';' || b == '(' || b == ')' || b == '[' || b == ']' || b == '{' || b == '}'
}

func isOperator(s string) bool {
	for i := 0; i < len(lxOPERATORS); i++ {
		if lxOPERATORS[i] == s {
			return true
		}
	}
	return false
}

func isKeyword(s string) bool {
	for i := 0; i < len(lxKEYWORDS); i++ {
		if lxKEYWORDS[i] == s {
			return true
		}
	}
	return false
}

func isIdStart(b byte) bool {
	return b == '_' || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isIdPart(b byte) bool {
	return isIdStart(b) || isNumStart(b)
}

func isNumStart(b byte) bool {
	return b >= '0' && b <= '9'
}

func isNumPart(b byte) bool {
	return isNumStart(b) || b == '.'
}

func (l *Lexer) next() bool {
	if l.cur == '\n' {
		l.lno++
	}
	if l.at < len(l.input) {
		l.cur = l.nxt
		l.nxt = l.char()
		return true
	}
	return false
}
func (l *Lexer) char() byte {
	if l.at > len(l.input) {
		return ' '
	}
	l.at++
	return l.input[l.at - 1]
}