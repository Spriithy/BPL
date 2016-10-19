package ast

import "github.com/Spriithy/BPL/compiler/parser"

type ASTNode struct {
	leaf bool
}

func (n *ASTNode) IsLeaf() bool {
	return n.leaf
}

// CONSTANT NODE ---------------------------------------------------------------

type ConstantNode struct {
	ASTNode
	Type  parser.ConstantType
	Value interface{}
}

func MakeConstantNode(t parser.ConstantType, val interface{}) *ConstantNode {
	return &ConstantNode{ASTNode{true}, t, val}
}
