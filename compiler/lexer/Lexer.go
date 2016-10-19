package lexer

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"strings"
)

var lxKeywords = map[string]bool{
	"const" : true, "static" : true, "public" : true, "private" : true,
	"native" : true, "if" : true, "else" : true, "unless" : true,
	"for" : true, "while" : true, "switch" : true, "case" : true,
	"break" : true, "continue" : true, "default" : true, "in" : true,
	"return" : true, "null" : true, "true" : true, "false" : true,
	"bool" : true, "byte" : true, "int" : true, "long" : true, "uint" : true,
	"ulong" : true, "float" : true, "double" : true, "complex" : true,
	"string" : true, "import" : true, "package" : true, "new" : true,
	"enum" : true,
}

var lxOperators = map[string]string{
	"=":"Assign", "+":"Plus", "-":"Minus", "*":"Times", "/":"Div", "%":"Mod",
	"~":"Not", "|":"bOr", "^":"bXor", "&":"bAnd", "<":"Lt", ">":"Gt",
	"!":"Bang", "++":"PlusPlus", "--":"MinusMinus", "**":"Pow", "->":"Arrow",
	"==":"Equals", "<<":"Lshift", ">>":"Rshift", "!=":"NotEquals",
	"+=":"OpAssign", "-=":"OpAssign", "*=":"OpAssign", "/=":"OpAssign",
	"**=":"OpAssign", "<<=":"OpAssign", ">>=":"OpAssign",
	"&=":"OpAssign", "|=":"OpAssign", "^=":"OpAssign", "~=":"OpAssign",
	"||":"Or", "&&":"And", "...":"Ellipsis", ".":"Dot", "::":"ColBlock",
	":":"Colon", ",":"Comma", "(":"LParen", ")":"RParen", "[":"LBrack", "]":"RBrack",
	"{":"LBrace", "}":"RBack",
}

/*
Builtin Token types :
---------------------
Unknown
EOF
Newline
Keyword
Identifier
Number
Character
String
LBrace
RBrace
LParen
RParen
LBrack
RBrack
Dot
Ellipsis
Comma
Colon
Semicolon
Assign
AssignOP
Arrow
ColBlock
Plus
Minus
Times
Divide
Mod
PlusPlus
MinusMinus
Pow
Lt
Leq
Lshift
NotEquals
Equals
Rshift
Geq
Gt
Bang
Not
Or
And
Xor
bAnd
bOr
 */

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
	return fmt.Sprintf("[%4d]\t%-14s " + vfmt, t.Lno, t.Kind, sym)
}

type lexer struct {
	path, input   string
	pos, caret    int
	lns, lne, lno int
	List, tail    *Token
}

func Lexer(path string) *lexer {
	lxr := new(lexer)
	lxr.path = path

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	lxr.input = string(content) + "\n\n"

	lxr.List = &Token{"EOF", "", Pos{0, 0}, Pos{0, 0}, 0, nil}
	lxr.tail = lxr.List

	return lxr
}

func (l *lexer) errorf(format string, a ...interface{}) {
	fmt.Printf("%s#%d: ", l.path, l.lno)
	fmt.Printf(format + "\n", a ...)
	os.Exit(1)
}

func (l *lexer) warningf(format string, a... interface{}) {
	fmt.Printf("%s#%d: ", l.path, l.lno)
	fmt.Printf(format + "\n", a ...)
}

func (l *lexer) logf(format string, a... interface{}) {
	fmt.Printf(format + "\n", a ...)
}

func (l *lexer) Lex() {
	l.lno, l.lns, l.lne, l.pos, l.caret = 1, 0, 0, 0, 0

	l.peekLine()
	for ; l.pos < len(l.input) - 2; l.pos++ {
		cc := l.input[l.pos]
		if l.input[l.pos:l.pos + 2] == "//" {
			for ; l.input[l.pos + 1] != '\n'; l.pos++ {}
			l.caret = l.pos
		} else if l.input[l.pos:l.pos + 2] == "/*" {
			for ; l.input[l.pos:l.pos + 2] != "*/"; l.pos++ {
				if l.input[l.pos] == '\n' {
					l.lno++
					l.lns = l.pos
					l.peekLine()
				}
			}
			l.pos++
			l.caret = l.pos
		} else if cc == ' ' || cc == '\t' || cc == '\r' {
			continue
		} else if isIdStart(cc) {
			l.caret = l.pos
			for ; isIdPart(l.input[l.caret]); l.caret++ {}

			if lxKeywords[l.input[l.pos:l.caret]] {
				l.addToken("Keyword")
			} else {
				l.addToken("Identifier")
			}
			l.pos = l.caret - 1
		} else if cc >= '0' && cc <= '9' {
			l.caret = l.pos
			d := false
			for ; isNumPart(l.input[l.caret]); l.caret++ {
				if l.input[l.caret + 1] == '.' && !d {
					d = true
				} else if l.input[l.caret + 1] == '.' {
					l.errorf("Malformed decimal number on line %d\n\t%s\n\t%s", l.lno, l.input[l.lns + 1:l.lne], strings.Repeat(" ", l.pos - l.lns - 1) + strings.Repeat("~", l.caret - l.pos + 1) + "^")
				}
			}
			l.addToken("Number")
			l.pos = l.caret - 1
		} else if cc == '\'' {
			chr := byte(0)
			if l.input[l.pos + 1] == '\\' {
				if l.input[l.pos + 3] != '\'' {
					l.errorf("Expected end of character litteral on line %d\n\t%s\n\t%s", l.lno, l.input[l.lns + 1:l.lne], strings.Repeat(" ", l.pos - l.lns - 1) + strings.Repeat("~", l.caret - l.pos + 4) + "^")
				}
				switch l.input[l.pos + 2] {
				case 'n': chr = '\n'
				case 't': chr = '\t'
				case 'r': chr = '\r'
				case '\'':chr = '\''
				case '\\':chr = '\\'
				default:
					l.errorf("Invalid escape sequence in character litteral on line %d\n\t%s\n\t%s", l.lno, l.input[l.lns + 1:l.lne], strings.Repeat(" ", l.pos - l.lns - 1) + strings.Repeat("~", l.caret - l.pos + 5))
				}
				l.pos += 3
			} else {
				chr = l.input[l.pos + 1]
				if l.input[l.pos + 2] != '\'' {
					l.errorf("Expected end of character litteral on line %d\n\t%s\n\t%s", l.lno, l.input[l.lns + 1:l.lne], strings.Repeat(" ", l.pos - l.lns - 1) + strings.Repeat("~", l.caret - l.pos + 3) + "^")
				}
				l.pos += 2
			}
			l.appendToken("Character", string(chr))
		} else if cc == '"' {
			esc := false
			str := ""
			l.caret = l.pos + 1
			for ; l.input[l.caret] != '"'; l.caret++ {
				if l.input[l.caret] == '\\' {
					esc = !esc
				}

				if esc {
					switch l.input[l.caret + 1] {
					case 'n': str += "\n"
					case 't': str += "\t"
					case 'r': str += "\r"
					case '"': str += "\""
					case '\\':str += "\\"
					default:
						l.errorf("Invalid escape sequence in string litteral on line %d\n\t%s\n\t%s", l.lno, l.input[l.lns + 1:l.lne], strings.Repeat(" ", l.pos - l.lns - 1) + strings.Repeat("~", l.caret - l.pos + 1) + "^")
					}
					esc = false
					l.caret++
				} else {
					str += string(l.input[l.caret])
				}
			}
			l.pos = l.caret
			l.appendToken("String", str)
		} else if lxOperators[string(l.input[l.pos])] != "" {
			l.caret = l.pos + 1
			for ; lxOperators[l.input[l.pos:l.caret]] != ""; l.caret++ {}
			l.caret--
			l.addToken(lxOperators[l.input[l.pos:l.caret]])
			l.pos = l.caret - 1
		} else if cc == '\n' {
			l.lno++
			l.lns = l.pos
			l.peekLine()
			if l.tail.Sym != "NewLine" {
				l.appendToken("--------------", "NewLine")
				continue
			}
		} else {
			print(string(cc))
		}
	}
	l.appendToken("~~~~~~~~~~~~~~", "EOF")
}

func (l *lexer) peekLine() {
	i := l.lns
	for ; l.input[i + 1] != '\n'; i++ {}
	l.lne = i + 1
}

func isIdStart(b byte) bool {
	return b == '_' || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func isIdPart(b byte) bool {
	return isIdStart(b) || b >= '0' && b <= '9'
}

func isNumPart(b byte) bool {
	return b >= '0' && b <= '9' || b == '.'
}

func (l *lexer) appendToken(kind, value string) {
	l.tail.Next = &Token{kind, value, Pos{l.pos, l.caret}, Pos{l.lns, l.lne}, l.lno, nil}
	l.tail = l.tail.Next
}

func (l *lexer) addToken(kind string) {
	l.appendToken(kind, l.input[l.pos:l.caret])
}
