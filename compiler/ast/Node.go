package ast

import (
	"github.com/Spriithy/BPL/compiler/token"
)

func MakeNode(val token.Token) *Node {
	val.Next = nil
	return &Node{val, 0, nil, nil}
}

func MakeParentNode(val token.Token, children... *Node) *Node {
	node := MakeNode(val)
	for _, v := range children {
		node.AddChild(v)
	}
	return node
}

type NStack []*Node

func (s NStack) Empty() bool {
	return len(s) == 0
}

func (s NStack) PeekTop() *Node {
	if s.Empty() {
		return nil
	}
	return s[len(s) - 1]
}

func (s *NStack) Push(n *Node) {
	(*s) = append((*s), n)
}

func (s *NStack) Pop() *Node {
	if s.Empty() {
		return nil
	}
	d := (*s)[len(*s) - 1]
	(*s) = (*s)[:len(*s) - 1]
	return d
}

type Node struct {
	Tok     token.Token
	count   int

	Child   *Node
	Sibling *Node
}

func (n *Node) AddChild(node *Node) {
	if node == nil {
		return
	}

	if n.IsLeaf() {
		n.Child = node
		return
	}

	n.count++
	node.Sibling = n.Child
	n.Child = node
}

func (n *Node) AddSibling(node *Node) {
	if node == nil {
		return
	}

	p := n
	for ; p.Sibling != nil; p = p.Sibling {}
	p.Sibling = node
}

func (n *Node) IsLeaf() bool {
	return n.Child == nil
}

func (n *Node) String() string {
	str := n.Print("")
	return str[:len(str) - 1]
}

func (n *Node) Print(prefix string) string {
	pat := "┣━ "
	if n.Sibling == nil {
		pat = "┗━ "
	}

	ext := "    "
	if n.Child != nil && n.Sibling != nil {
		ext = "┃   "
	}

	str := prefix + pat + n.Tok.ShortString() + "\n"
	if n.Child != nil {
		str += n.Child.Print(prefix + ext)
	}

	if n.Sibling != nil {
		str += n.Sibling.Print(prefix)
	}

	return str
}
