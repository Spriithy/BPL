package parser

import (
	"github.com/Spriithy/BPL/compiler/token"
	"fmt"
	"os"
)

var prOps = map[string]struct {
	prec   int
	rAssoc bool
}{
	"++" : {30, false}, "--": {30, false},
	"!"  : {29, true}, "~"  : {30, true},
	"**" : {28, true},
	"*"  : {27, false}, "/"  : {27, false}, "%" : {27, false},
	"+"  : {26, false}, "-"  : {26, false},
	">>" : {25, false}, "<<" : {25, false},
	">"  : {24, false}, ">=" : {24, false}, "<" : {24, false}, "<=" : {24, false},
	"==" : {23, false}, "!=" : {23, false},
	"&"  : {22, false},
	"^"  : {21, false},
	"|"  : {20, false},
	"&&" : {19, false},
	"||" : {18, false},
	"="  : {10, true}, "+="  : {10, true}, "-=" : {10, true}, "*="  : {10, true},
	"/=" : {10, true}, "**=" : {10, true}, "^=" : {10, true}, "~="  : {10, true},
	"|=" : {10, true}, "&="  : {10, true}, "%=" : {10, true}, "<<=" : {10, true},
	">>=": {10, true},
	","  : {9, false},
}

type parser struct {
	path, src string
	lno       int
	tokens    token.TQueue
}

func Parser(path, source string, tokens token.TQueue) *parser {
	p := new(parser)
	p.path = path
	p.src = source
	p.tokens = tokens
	p.lno = p.tokens.PeekHead().Lno
	return p
}

func (p *parser) errorf(format string, a ...interface{}) {
	fmt.Printf("%s#%d: ", p.path, p.lno)
	fmt.Printf(format + "\n", a ...)
	os.Exit(1)
}

func (p *parser) warningf(format string, a... interface{}) {
	fmt.Printf("%s#%d: ", p.path, p.lno)
	fmt.Printf(format + "\n", a ...)
}

func (p *parser) logf(format string, a... interface{}) {
	fmt.Printf(format + "\n", a ...)
}

func (p *parser) Parse() {
	rpn := p.ShuntingYard(p.tokens)

	ptr := rpn.PeekHead()
	for ; ptr != nil; ptr = ptr.Next {
		fmt.Println(ptr.String())
	}
}

/*
Shunting Yard Implementation to parse Expressions
-------------------------------------------------

While there are tokens to be read:
	Read a token.
	If the token is a number, then push it to the output queue.
	If the token is a function token, then push it onto the stack.
	If the token is a function argument separator (e.g., a comma):
		Until the token at the top of the stack is a left parenthesis, pop operators off the stack onto the output queue.
		(* If no left parentheses are encountered, either the separator was misplaced or parentheses were mismatched. *)
	If the token is an operator, o1, then:
		while there is an operator token o2, at the top of the operator stack and either
			o1 is left-associative and its precedence is less than or equal to that of o2, or
			o1 is right associative, and has precedence less than that of o2,
				pop o2 off the operator stack, onto the output queue;
		at the end of iteration push o1 onto the operator stack.
	If the token is a left parenthesis (i.e. "("), then push it onto the stack.
	If the token is a right parenthesis (i.e. ")"):
		Until the token at the top of the stack is a left parenthesis, pop operators off the stack onto the output queue.
		Pop the left parenthesis from the stack, but not onto the output queue.
		If the token at the top of the stack is a function token, pop it onto the output queue.
		If the stack runs out without finding a left parenthesis, then there are mismatched parentheses.
	When there are no more tokens to read:
		While there are still operator tokens in the stack:
			If the operator token on the top of the stack is a parenthesis, then there are mismatched parentheses.
			Pop the operator onto the output queue.
Exit.
 */

func (p *parser) ShuntingYard(input token.TQueue) token.TQueue {
	stack := token.TokenStack()
	output := token.TokenQueue()

	for tok := input.Dequeue(); tok.Sym != "EOF"; tok = input.Dequeue() {
		switch tok.Kind {
		case "LParen":
			stack.Push(tok) // push "(" to stack
		case "RParen":
			for {
				// pop item ("(" or operator) from stack
				if stack.Empty() {
					p.errorf("Unmatched parenthesis on line %d, expected '(' to match closing parenthesis in expression", p.lno)
				}

				op := stack.Pop()
				if op.Sym == "(" {
					break // discard "("
				}
				output.Enqueue(op)
			}
		case "Identifier":

		default:
			if o1, isOp := prOps[tok.Sym]; isOp {
				// token is an operator
				for !stack.Empty() {
					// consider top item on stack
					op := stack.PeekTop()
					if o2, isOp := prOps[op.Sym]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					output.Enqueue(stack.Pop()) // pop it and add it to result
				}
				// push operator (the new one) to stack
				stack.Push(tok)
			} else {
				output.Enqueue(tok)
			}
		}
	}

	// drain stack to result
	for !stack.Empty() {
		if stack.PeekTop().Sym == "(" {
			p.errorf("Unmatched parenthesis on line %d, expected ')' to match previous parenthesis in expression", p.lno)
		}
		output.Enqueue(stack.Pop())
	}

	return *output
}