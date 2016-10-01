package vm

import (
	"fmt"
)

// The NativeType enumerates all the types that are built-in for the VM
type NativeType int

const (
	// Void represents the absence of type but not value. An empty value would
	// be filled with a nil value
	VoidType NativeType = iota

	// Bool type is either True or False
	// Bools are Comparable
	BoolType

	// Byte is used to represent unsigned integers up to 255, they are used
	// to manipulate characters most of the time.
	// Bytes are Comparable
	ByteType

	// Int is the default number representation in the VM architecture and
	// are used for general purpose integer calculations
	// Ints are Comparable
	IntType

	// UInt define the unsigned int type used in size calculations.
	// Computations between UInts and Ints always result in UInts rounded
	// at Zero if the result were to be a negative number
	// UInts are comparable
	UIntType

	// TODO
	LongType

	// TODO
	ULongType

	// Single Precison decimal numbers are represented with the Float type
	// Floats are Comparable
	FloatType

	// Double Precison decimal numbers are represented with the Double type
	// Doubles are Comparable
	DoubleType

	// Complex numbers are represented using two Double components each respectively
	// representing the real and imaginary part of the Complex number
	// Complex numbers are Comparable
	ComplexType

	// Strings are nothing more than Immutable Arrays of Bytes
	// Strings are Comparable, Sliceable and Partialy Indexable (read-only)
	StringType

	// An Object is a collection of the above types, gathered in fields
	// Objects are not Comparable but are Accessible
	ObjectType
) // NativeTypes end


// Declares Golang builtin types aliases for the VM system

func TypeNameOf(o NativeObject) string {
	switch o.TypeOf() {
	case VoidType:
		return "Void";
	case BoolType:
		return "Bool";
	case ByteType:
		return "Byte";
	case IntType:
		return "Int";
	case UIntType:
		return "UInt";
	case FloatType:
		return "Float";
	case DoubleType:
		return "Double";
	case ComplexType:
		return "Complex";
	case StringType:
		return "String"
	case ObjectType:
		return "Object"
	default:
		return "<" + o.String() + ">"
	}
}

// Native Interfaces describe the possible behaviors of Objects at the VM Runtime
type NativeObject interface {
	TypeOf() NativeType
	String() string
}

type Comparable interface {
	Equals(other Comparable) bool
}

type Indexable interface {
	GetItem(at int) NativeObject
	SetItem(at int, o NativeObject)
}

type Sliceable interface {
	Slice(start, end int) *Sliceable
}

type Accessible interface {
	GetField(name int) NativeObject
	SetField(name int, o NativeObject)
}
// Native Interfaces end

type NativeValue struct {
	typ     NativeType

	string  *StringObject
	bool    bool
	byte    byte
	int     int
	uint    uint
	long    int64
	ulong   uint64
	float   float32
	double  float64
	complex complex128
}

func NewNative(of NativeType) NativeObject {
	switch of {
	case VoidType:
		return NewVoid()
	case BoolType:
		return NewBool(false)
	case ByteType:
		return NewByte(0)
	case IntType:
		return NewInt(0)
	case UIntType:
		return NewUInt(0)
	case LongType:
		return NewLong(0)
	case ULongType:
		return NewULong(0)
	case FloatType:
		return NewFloat(0.0)
	case DoubleType:
		return NewDouble(0.0)
	case ComplexType:
		return NewComplex(complex(0.0, 0.0))
	case StringType:
		return NewStringObject("")
	}
	return NewObject()
}

func NewObject() NativeObject {
	return nil
}

func NewVoid() *NativeValue {
	return &NativeValue{typ: VoidType}
}

func NewBool(b bool) *NativeValue {
	return &NativeValue{typ: BoolType, bool: b}
}

func NewByte(b byte) *NativeValue {
	return &NativeValue{typ: ByteType, byte: b}
}

func NewInt(i int) *NativeValue {
	return &NativeValue{typ: IntType, int: i}
}

func NewUInt(ui uint) *NativeValue {
	return &NativeValue{typ: UIntType, uint: ui}
}

func NewLong(l int64) *NativeValue {
	return &NativeValue{typ: LongType, long: l}
}

func NewULong(ul uint64) *NativeValue {
	return &NativeValue{typ: ULongType, ulong: ul}
}

func NewFloat(f float32) *NativeValue {
	return &NativeValue{typ: FloatType, float: f}
}

func NewDouble(d float64) *NativeValue {
	return &NativeValue{typ: DoubleType, double: d}
}

func NewComplex(c complex128) *NativeValue {
	return &NativeValue{typ: ComplexType, complex: c}
}

func NewString(s string) *NativeValue {
	return &NativeValue{typ: StringType, string: NewStringObject(s)}
}

func (o *NativeValue) Equals(other NativeValue) bool {
	if o.typ != other.typ {
		return false
	}

	switch o.typ {
	case VoidType:
		return true
	case BoolType:
		return o.bool == other.bool
	case ByteType:
		return o.byte == other.byte
	case IntType:
		return o.int == other.int
	case UIntType:
		return o.uint == other.uint
	case LongType:
		return o.long == other.long
	case ULongType:
		return o.ulong == other.ulong
	case FloatType:
		return o.float == other.float
	case DoubleType:
		return o.double == other.double;
	case ComplexType:
		return o.complex == other.complex
	case StringType:
		return o.string.val == other.string.val
	}

	return false
}

func (o *NativeValue) TypeOf() NativeType {
	return o.typ
}

func (o *NativeValue) String() string {
	switch o.typ {
	case VoidType:
		return "Void";
	case BoolType:
		return fmt.Sprintf("%v", o.bool)
	case ByteType:
		return fmt.Sprintf("%#v", o.byte)
	case IntType:
		return fmt.Sprintf("%v", o.int)
	case UIntType:
		return fmt.Sprintf("%v", o.uint)
	case LongType:
		return fmt.Sprintf("%v", o.long)
	case ULongType:
		return fmt.Sprintf("%v", o.ulong)
	case FloatType:
		return fmt.Sprintf("%v", o.float)
	case DoubleType:
		return fmt.Sprintf("%v", o.double)
	case ComplexType:
		return fmt.Sprintf("%v", o.complex)
	case StringType:
		return o.string.Repr()
	}
	return "Unkown"
}

func (o *NativeValue) Bool() bool {
	return o.bool
}

func (o *NativeValue) Byte() byte {
	return o.byte
}

func (o *NativeValue) Int() int {
	return o.int
}

func (o *NativeValue) UInt() uint {
	return o.uint
}

func (o *NativeValue) Long() int64 {
	return o.long
}

func (o *NativeValue) ULong() uint64 {
	return o.ulong
}

func (o *NativeValue) Float() float32 {
	return o.float
}

func (o *NativeValue) Double() float64 {
	return o.double
}

func (o *NativeValue) Complex() complex128 {
	return o.complex
}

