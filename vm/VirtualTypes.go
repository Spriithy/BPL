package vm

import (
	"fmt"
	"math"
)

type VirtualType int

const (
	VIRTUAL_NULL VirtualType = iota
	VIRTUAL_INTEGER
	VIRTUAL_REAL
)

const (
	VIRTUAL_TRUE VirtualInt = VirtualInt(1)
	VIRTUAL_FALSE VirtualInt = VirtualInt(0)
)

type VirtualInt  int64
type VirtualReal float64
type VirtualReference uint16
type VirtualNull VirtualInt

type VirtualValue interface {
	Type() VirtualType  // Returns the actual type of the data stored in there

	ToInt() int64       // Returns int value if VirtValue is VirtInt, or int part if VirtReal
	ToReal() float64    // Returns real value if VirtValue is VirtReal, or float cast if VirtInt

	String() string     // Returns the representation of the Virtual Value
}

func (n VirtualNull) Type() VirtualType {
	return VIRTUAL_NULL
}

func (n VirtualNull) String() string {
	return "null"
}

func (n VirtualNull) ToInt() int64 {
	return -1
}

func (n VirtualNull) ToReal() float64 {
	return math.NaN()
}

// VIRTUAL INTEGER REPRESENTATION ----------------------------------------------

func (i VirtualInt) ToInt() int64 {
	return int64(i) // Back to native type
}

func (i VirtualInt) ToReal() float64 {
	return float64(i) // Simple int->float convertion
}

func (i VirtualInt) Type() VirtualType {
	return VIRTUAL_INTEGER
}

func (i VirtualInt) String() string {
	return fmt.Sprintf("%d", i)
}

// VIRTUAL REAL REPRESENTATION -------------------------------------------------

func (r VirtualReal) ToInt() int64 {
	return int64(r) // Int part (== Round Down)
}

func (r VirtualReal) ToReal() float64 {
	return float64(r) // Back to native type
}

func (r VirtualReal) Type() VirtualType {
	return VIRTUAL_REAL
}

func (r VirtualReal) String() string {
	return fmt.Sprintf("%v", float64(r))
}
