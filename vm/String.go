package vm

import (
	"fmt"
	"log"
)

type StringObject struct {
	val string
	len int
}

func NewStringObject(of string) *StringObject {
	return &StringObject{val: of, len: len(of)}
}

func (o *StringObject) TypeOf() NativeType {
	return StringType
}

func (o *StringObject) String() string {
	return o.val
}

func (o *StringObject) Repr() string {
	return fmt.Sprintf("%#v", o.val)
}

func (o *StringObject) Equals(other StringObject) bool {
	return o.val == other.val;
}

func (o *StringObject) ByteAt(at int) byte {
	return o.val[at]
}

func (o *StringObject) Slice(start, end int) *StringObject {
	if start >= o.len || end > o.len {
		log.Fatalf("VM-Error: slice index out of range:\n\ttrying to apply slice(%d, %d) when source String is %d-Bytes long", start, end, o.len)
	}

	return &StringObject{val: o.val[start:end], len: end - start}
}

func (o *StringObject) Concat(with NativeObject) *StringObject {
	return NewStringObject(string(o.val + with.String()))
}
