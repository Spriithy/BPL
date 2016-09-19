package vm

import (
	"fmt"
	"os"
	"log"
)

type RTType int

const (
	None RTType = iota

	Byte
	Int
	UInt
	Float

	String
	Array
)

type RTObject interface {
	Type() RTType
	RTTypeStr() string
	RTValue() interface{}
}

type RTNative struct {
	value interface{}
	t     RTType
}

type RTArray struct {
	size int
	data []RTObject
}

type RTString struct {
	size  int
	value string
}

/*
** RTObject constructors implementations
*/

func InitNone() *RTNative {
	n := new(RTNative)
	n.value = nil
	n.t = None
	return n
}

func InitNative(t RTType, v interface{}) *RTNative {
	n := new(RTNative)
	n.value = v
	n.t = t
	return n
}

func InitArray(size int) *RTArray {
	a := new(RTArray)
	a.size = size
	a.data = make([]RTObject, size)

	for i, _ := range a.data {
		a.data[i] = InitNone()
	}

	return a
}

func InitString(str string) *RTString {
	s := new(RTString)
	s.value = str
	s.size = len(str)
	return s
}

func NativeByte(b uint8) *RTNative {
	return InitNative(Byte, b)
}

func NativeInt(i int64) *RTNative {
	return InitNative(Int, i)
}

func NativeUInt(i uint64) *RTNative {
	return InitNative(UInt, i)
}

func NativeFloat(f float64) *RTNative {
	return InitNative(Float, f)
}

/*
** RTCopy() functions implementations
*/

func (rtn *RTNative) RTCopy() *RTNative {
	return InitNative(rtn.t, rtn.value)
}

func (rta *RTArray) RTCopy() *RTArray {
	a := InitArray(rta.size)
	a.data = rta.data
	return a
}

func (rts *RTString) RTCopy() *RTString {
	return InitString(rts.value)
}

/*
** RTTypeStr () functions implementations
*/

func (rtn RTNative) RTTypeStr() string {
	switch rtn.t {
	case Byte: return "Byte"
	case Int: return "Int"
	case UInt: return "UInt"
	case Float: return "Float"
	}
	return "None"
}

func (rtn RTNative) Type() RTType {
	return rtn.t
}

func (rta RTArray) RTTypeStr() string {
	return fmt.Sprintf("Array[%d]", rta.size)
}

func (rtn RTArray) Type() RTType {
	return Array
}

func (rts RTString) RTTypeStr() string {
	return "String"
}

func (rts RTString) Type() RTType {
	return String
}

/*
** RTValue () functions implementations
*/

func (rtn RTNative) RTValue() interface{} {
	return rtn.value
}

func (rta RTArray) RTValue() interface{} {
	return rta.data
}

func (rts RTString) RTValue() interface{} {
	return fmt.Sprintf("%#v", rts.value)
}


/*
** RTArray related functions
*/

func (rta *RTArray) SetItem(at int, item RTObject) {
	if at >= rta.size {
		fmt.Errorf("VM-Error: Assignement out of Array bounds. (index = %d, size = %d) ", at, rta.size)
		os.Exit(1)
	}
	rta.data[at] = item
}

func (rta *RTArray) GetItem(at int) RTObject {
	if at >= rta.size {
		fmt.Errorf("VM-Error: Indexing out of Array bounds. (index = %d, size = %d)", at, rta.size)
		os.Exit(1)
	}
	return rta.data[at]
}

/*
** RTString related functions
*/

func (rts *RTString) SetByteAt(at, char int) {
	if at >= rts.size {
		log.Fatal(fmt.Errorf("VM-Error: Assignement out of String bounds. (index = %d, size = %d) ", at, rts.size))
	}
	rts.value = rts.value[:at] + string(char) + rts.value[at + 1:]
}

func (rts *RTString) GetByte(at int) RTNative {
	if at >= rts.size {
		log.Fatal(fmt.Errorf("VM-Error: Indexing out of String bounds. (index = %d, size = %d)", at, rts.size))
	}
	return *InitNative(Byte, rts.value[at])
}
