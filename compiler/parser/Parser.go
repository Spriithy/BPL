package parser

import (
	"github.com/Spriithy/BPL/compiler/ast"
)

type ConstantType int

const (
	CT_UNKNOWN ConstantType = iota
	CT_NULL
	CT_INTEGER
	CT_REAL
	CT_BOOL
	CT_CHAR
	CT_STRING
)

type parser struct {
	ast.ASTNode
}

func Parser() *parser {
	return new(parser)
}