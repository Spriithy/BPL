package compiler

import (
	"io/ioutil"
	"fmt"
	"log"
	"errors"
)

var lxOPERATORS = []string{
	"=", "+", "-", "*", "/", "%", "~", "|", "^", "&", "<", ">", "!", "++",
	"--", "**", "//", "==", "<<", ">>", "!=", "+=", "-=", "*=", "/=", "**=",
	"//=", "<<=", ">>=", "&=", "|=", "^=", "~=", "||", "&&", "...", ".",
	"::", "->" }

var lxKEYWORDS = []string{
	"const", "struct", "enum", "public", "private", "static", "new", "if",
	"else", "for", "while", "unless", "return", "switch", "match", "true",
	"false", "null", "final", "this", "break", "continue", "default", "int",
	"long", "float", "double", "complex", "byte", "string", "import", "uint",
	"ulong", "bool", "native", "package" }

type Lexer struct {
	Path, input  string
	at, lno      int
	cur, nxt     byte
	ws           bool
	Tokens, last *Token
}

func NewLexer(path string) *Lexer {
	l := new(Lexer)
	l.Path = path
	l.ws = false
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

func (l *Lexer) addtoken(tok *Token) {
	l.last.Next = tok
	l.last = l.last.Next
}

func (l *Lexer) process() {
	fmt.Printf("%-5d: %#v\n", l.at, string(l.cur))
	c, n := string(l.cur), string(l.nxt)
	if c + n == "//" {
		l.comment(0)
	} else if c + n == "/*" {
		l.comment(1)
	} else if c == "\n" {
		if !l.ws && l.last.typ != tNL {
			l.addtoken(NL(l.lno))
			l.ws = true
			return
		}
	} else if c == "\t" || c == " " || c == "\r" {
		return
	} else if c == "\"" {
		l.dostring()
	} else if c == "'" {
		l.dochar()
	} else if isPunctuation(l.cur) {
		l.punctuation()
	} else if isIdStart(l.cur) {
		l.identifier()
	} else if isNumStart(l.cur) {
		l.number()
	} else if isOperator(c) {
		l.operator()
	} else {
		log.Fatal(errors.New("lex: Unknown lexing state on line " + string(l.lno)))
	}
	l.ws = false
}

func (l *Lexer) punctuation() {
	l.addtoken(Punctuation(l.cur, l.lno))
}

func (l *Lexer) operator() {
	op1 := string(l.cur)
	op2 := op1 + string(l.nxt)
	l.next()
	op3 := op2 + string(l.nxt)
	if isOperator(op3) {
		l.next()
		l.addtoken(Operator(op3, l.lno))
		return
	}

	if isOperator(op2) {
		l.addtoken(Operator(op2, l.lno))
		return
	}

	if isOperator(op1) {
		l.addtoken(Operator(op1, l.lno))
		l.at--
		l.nxt = l.cur
		return
	}
}

func (l *Lexer) dochar() {
	chr := byte(0)
	n := l.nxt

	if n == '\'' || n == '\n' || n == '\t' {
		log.Fatal(fmt.Errorf("lex: Invalid character (empty ?) literal on line %d", l.lno))
	}

	l.next()
	chr = l.cur
	if chr == '\\' {
		switch l.nxt {
		case '\\': break
		case 'n' : chr = '\n'
		case 't' : chr = '\t'
		case 'r' : chr = '\r'
		case '\'': chr = '\''
		default:
			log.Fatal(fmt.Errorf("lex: Malformed escape sequence '\\%s' on line %d", string(l.nxt), l.lno))
			return
		}
		l.next()
	}

	if l.nxt != '\'' {
		log.Fatal(fmt.Errorf("lex: Invalid multi-character literal on line %d. Use '\"' to specify String literals", l.lno))
	}

	l.next()
	l.addtoken(Char(chr, l.lno))
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
				log.Fatal(fmt.Errorf("lex: Malformed escape sequence '\\%s' on line %d", string(l.nxt), l.lno))
			}
			l.next()
			escape = false
		} else {
			str += string(l.cur)
		}
		l.next()
	}
	l.addtoken(String(str, l.lno))
}

func (l *Lexer) identifier() {
	id := string(l.cur)
	for ; isIdPart(l.nxt); {
		id += string(l.nxt)
		l.next()
	}

	if isKeyword(id) {
		l.addtoken(Keyword(id, l.lno))
	} else {
		l.addtoken(Identifier(id, l.lno))
	}
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
	l.addtoken(Number(num, l.lno))
}

func isPunctuation(b byte) bool {
	return b == ',' || b == ':' || b == ';' || b == '(' || b == ')' || b == '[' || b == ']' || b == '{' || b == '}'
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
		return '\n'
	}
	l.at++
	return l.input[l.at - 1]
}