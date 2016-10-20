package token

func NilTok() *Token {
	return &Token{"", "nil", Pos{0, 0}, Pos{0, 0}, 0, nil}
}

// STACK -----------------------------------------------------------------------

type TStack struct {
	tos *Token
}

func TokenStack() *TStack {
	return &TStack{NilTok()}
}

func (s *TStack) Push(tok *Token) {
	tok.Next = s.tos.Next
	s.tos.Next = tok
}

func (s *TStack) Pop() *Token {
	tok := s.tos.Next
	s.tos.Next = tok.Next
	tok.Next = nil
	return tok
}

func (s *TStack) PeekTop() *Token {
	return s.tos.Next
}

func (s *TStack) Empty() bool {
	return s.Size() == 0
}

func (s *TStack) Size() int {
	i := 0
	p := s.PeekTop()
	for ; p != nil; p = p.Next {
		i++
	}
	return i
}

func (s *TStack) String() string {
	str := ""
	p := s.PeekTop()
	for ; p != nil; p = p.Next {
		str += p.Sym + ", "
	}

	if len(str) > 0 {
		return "t[" + str[:len(str) - 2] + "]"
	}
	return "t[]"
}

// QUEUE -----------------------------------------------------------------------

type TQueue struct {
	head, tail *Token
}

func TokenQueue() *TQueue {
	q := new(TQueue)
	q.head = NilTok()
	q.head.Sym = "QUEUE"
	q.tail = q.head
	return q
}

func (q *TQueue) Enqueue(tok *Token) {
	q.tail.Next = tok
	tok.Next = nil
	q.tail = tok
}

func (q *TQueue) Dequeue() *Token {
	tok := q.head.Next
	q.head.Next = tok.Next
	tok.Next = nil
	return tok
}

func (q *TQueue) PeekTail() *Token {
	return q.tail
}

func (q *TQueue) PeekHead() *Token {
	return q.head.Next
}

