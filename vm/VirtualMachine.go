package vm

import (
	"fmt"
	"unicode/utf8"
	"math"
)

type vm struct {
	Stack
	ErrorHandler

	ipool  []int64
	rpool  []float64
	refs   []VirtualReference

	code   []Bytecode
	ip, fp int
}

func VirtualMachine(code[]Bytecode) *vm {
	this := &vm{
		Stack:      *NewStack(),
		ipool:      make([]int64, 0),
		rpool:      make([]float64, 0),
		refs:       make([]VirtualReference, 0),
		code:       code,
		ip:         0,
		fp:         0,
	}
	this.ErrorHandler = *NewVirtualErrorHandler(this)
	return this
}

func (v *vm) Start() {
	v.run(0)
}


// initial instruction pointer
func (v *vm) run(iip int) {
	v.sp = -1
	v.ip = iip

	for {
		v.disassemble()

		// Fetch the instruction to execute
		op := v.code[v.ip]
		v.ip++

		switch op {
		case HALT:
			v.cleanup()
			return
		case NULLREF:
			v.Push(VirtualNull(0))
		// IO START
		case LNFEED: println()
		case STDOUT_VAL: print(v.Pop().String())
		case STDOUT_UNI:
			v1 := v.Pop()
			if v1.Type() != VIRTUAL_INTEGER {
				v.stress(
					INTEGER_VALUE_EXPECTED_ERROR,
					"Cannot print non-integer value as Unicode character!")
			}
			buf := make([]byte, 8)
			utf8.EncodeRune(buf, rune(v1.ToInt()))
			print(string(buf))
		case STDOUT_REF:
			val := v.Pop()
			if val.Type() != VIRTUAL_INTEGER {
				v.stress(
					INVALID_VIRTUAL_REFERENCE_PROCESS_ERROR,
					"Processing non-integer reference!")
			}
			fmt.Printf("*(0x%X)", val.ToInt())
		case LNOUT_VAL: print(v.Pop().String() + "\n")
		case LNOUT_UNI:
			v1 := v.Pop()
			if v1.Type() != VIRTUAL_INTEGER {
				v.stress(
					INTEGER_VALUE_EXPECTED_ERROR,
					"Cannot print non-integer value as Unicode character!")
			}
			buf := make([]byte, 8)
			utf8.EncodeRune(buf, rune(v1.ToInt()))
			print(string(buf) + "\n")
		case LNOUT_REF:
			val := v.Pop()
			if val.Type() != VIRTUAL_INTEGER {
				v.stress(
					INVALID_VIRTUAL_REFERENCE_PROCESS_ERROR,
					"Processing non-integer reference!")
			}
			fmt.Printf("*(0x%X)\n", val.ToInt())
		// START MEMORY
		case ICONST_0: v.PushI(0)
		case ICONST_1: v.PushI(1)
		case ICONST_2: v.PushI(2)
		case ICONST_3: v.PushI(3)
		case ICONST_4: v.PushI(4)
		case ICONST_5: v.PushI(5)
		case ICONST_N:
			v1 := int64(v.code[v.ip]) // TODO create pool
			v.ip++
			v.PushI(v1)
		case RCONST_0: v.PushR(0.0)
		case RCONST_1: v.PushR(1.0)
		case RCONST_N:
			v1 := v.rpool[v.ip + 1]
			v.ip++
			v.PushR(v1)
		case RCONST_E: v.PushR(math.E)
		case RCONST_PHI: v.PushR(math.Phi)
		case RCONST_PI: v.PushR(math.Pi)
		// END MEMORY
		case CMP:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 > v2 {
				v.PushI(1)
			} else if v1 < v2 {
				v.PushI(-1)
			} else {
				v.PushI(0)
			}
		case EQ:
			if v.PopR() == v.PopR() {
				v.Push(VIRTUAL_TRUE)
			} else {
				v.Push(VIRTUAL_FALSE)
			}
		case NEQ:
			if v.PopR() != v.PopR() {
				v.Push(VIRTUAL_TRUE)
			} else {
				v.Push(VIRTUAL_FALSE)
			}
		case LT:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 < v2 {
				v.Push(VIRTUAL_TRUE)
			} else {
				v.Push(VIRTUAL_FALSE)
			}
		case LEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 <= v2 {
				v.Push(VIRTUAL_TRUE)
			} else {
				v.Push(VIRTUAL_FALSE)
			}
		case GT:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 > v2 {
				v.Push(VIRTUAL_TRUE)
			} else {
				v.Push(VIRTUAL_FALSE)
			}
		case GEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 >= v2 {
				v.Push(VIRTUAL_TRUE)
			} else {
				v.Push(VIRTUAL_FALSE)
			}
		case NEG:
			v1 := v.Pop()
			switch v1.Type() {
			case VIRTUAL_INTEGER:
				v.PushI(-v1.ToInt())
			case VIRTUAL_REAL:
				v.PushR(-v1.ToReal())
			default:
				break
			}
		case INC:
			val := v.Pop()
			switch val.Type() {
			case VIRTUAL_INTEGER:
				v.PushI(val.ToInt() + 1)
			case VIRTUAL_REAL:
				v.PushR(val.ToReal() + 1.0)
			default:
				break
			}
		case DEC:
			v1 := v.Pop()
			switch v1.Type() {
			case VIRTUAL_INTEGER:
				v.PushI(v1.ToInt() - 1)
			case VIRTUAL_REAL:
				v.PushR(v1.ToReal() - 1.0)
			default:
				break
			}
		case DIV:
			v2 := v.PopR()
			v1 := v.PopR()
			if v2 == 0 {
				v.stress(
					ZERO_DIVISION_ERROR,
					"Cannot process division by 0, result is NaN")
			}
			v.PushR(v1 / v2)
		// Int related
		case IADD:
			v2 := v.PopI()
			v1 := v.PopI()
			v.PushI(v1 + v2)
		case ISUB:
			v2 := v.PopI()
			v1 := v.PopI()
			v.PushI(v1 - v2)
		case IMUL:
			v2 := v.PopI()
			v1 := v.PopI()
			v.PushI(v1 * v2)
		case IMOD:
			v2 := v.PopI()
			v1 := v.PopI()
			v.PushI(v1 % v2)
		case ISHR:
			v2 := v.PopI()
			v1 := v.PopI()

			if v2 < 0 {
				v.stress(
					UNSIGNED_INTEGER_EXPECTED_ERROR,
					"Shift count must be a positive integer!")
			}

			v.PushI(v1 >> uint64(v2))
		case ISHL:
			v2 := v.PopI()
			v1 := v.PopI()

			if v2 < 0 {
				v.stress(
					UNSIGNED_INTEGER_EXPECTED_ERROR,
					"Shift count must be a positive integer!")
			}

			v.PushI(v1 << uint64(v2))
		case IAND:
			v2 := v.PopI()
			v1 := v.PopI()
			v.PushI(v1 & v2)
		case IOR:
			v2 := v.PopI()
			v1 := v.PopI()
			v.PushI(v1 | v2)
		case IXOR:
			v2 := v.PopI()
			v1 := v.PopI()
			v.PushI(v1 ^ v2)
		case INOT:
			v1 := v.PopI()
			if v1 == 0 {
				v.Push(VIRTUAL_TRUE)
			} else {
				v.Push(VIRTUAL_FALSE)
			}
		case ICOMPL1:
			v.PushI(^v.PopI() - 1)
		case ICOMPL2:
			v.PushI(^v.PopI())
		// Real related
		case RADD:
			v2 := v.PopR()
			v1 := v.PopR()
			v.PushR(v1 + v2)
		case RSUB:
			v2 := v.PopR()
			v1 := v.PopR()
			v.PushR(v1 - v2)
		case RMUL:
			v2 := v.PopR()
			v1 := v.PopR()
			v.PushR(v1 * v2)
		case RSHL:
			v2 := v.PopI()
			v1 := v.PopR()
			v.PushR(v1 * math.Pow10(int(v2)))
		case RSHR:
			v2 := v.PopI()
			v1 := v.PopR()
			v.PushR(v1 * math.Pow10(-int(v2)))

		// TECHNICAL START
		case NOP: continue
		case DROP: v.Pop()
		case DUP:
			v1 := v.Pop()
			v.Push(v1); v.Push(v1)
		case DUP2:
			v1 := v.Pop()
			v2 := v.Pop()
			v.Push(v2); v.Push(v1)
			v.Push(v2); v.Push(v1)
		case SWAP:
			v1 := v.Pop()
			v2 := v.Pop()
			v.Push(v2); v.Push(v1)
		case SWAP2:
			v1 := v.Pop()
			v2 := v.Pop()
			v3 := v.Pop()
			v4 := v.Pop()
			v.Push(v2); v.Push(v1)
			v.Push(v4); v.Push(v3)
		// Branch if
		case BR:
			v.ip++
			v.ip = int(v.code[v.ip])
		case BR_0:
			v.ip++
			if v.PopI() == 0 {
				v.ip = int(v.code[v.ip])
			}
		case BR_N0:
			v.ip++
			if v.PopI() != 0 {
				v.ip = int(v.code[v.ip])
			}
		case BR_LT:
			v2 := v.PopR()
			v1 := v.PopR()
			v.ip++
			if v1 < v2 {
				v.ip = int(v.code[v.ip])
			}
		case BR_GT:
			v2 := v.PopR()
			v1 := v.PopR()
			v.ip++
			if v1 > v2 {
				v.ip = int(v.code[v.ip])
			}
		case BR_LEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			v.ip++
			if v1 <= v2 {
				v.ip = int(v.code[v.ip])
			}
		case BR_GEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			v.ip++
			if v1 >= v2 {
				v.ip = int(v.code[v.ip])
			}
		case BR_EQ:
			v2 := v.PopR()
			v1 := v.PopR()
			v.ip++
			if v1 == v2 {
				v.ip = int(v.code[v.ip])
			}
		case BR_NEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			v.ip++
			if v1 != v2 {
				v.ip = int(v.code[v.ip])
			}
		case BR_NUL:
			v1 := v.Pop()
			v.ip++
			if v1.Type() == VIRTUAL_NULL {
				v.ip = int(v.code[v.ip])
			}
		case BR_NNUL:
			v1 := v.Pop()
			v.ip++
			if v1.Type() != VIRTUAL_NULL {
				v.ip = int(v.code[v.ip])
			}
		// If-Else
		case IF_0:
			if v.PopI() == 0 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_N0:
			if v.PopI() != 0 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_LT:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 < v2 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_GT:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 > v2 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_LEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 <= v2 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_GEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 >= v2 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_EQ:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 == v2 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_NEQ:
			v2 := v.PopR()
			v1 := v.PopR()
			if v1 != v2 {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_NUL:
			v1 := v.Pop()
			if v1.Type() == VIRTUAL_NULL {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		case IF_NNUL:
			v1 := v.Pop()
			if v1.Type() != VIRTUAL_NULL {
				v.ip = int(v.code[v.ip + 1])
			} else {
				v.ip = int(v.code[v.ip + 2])
			}
		// TECHNICAL END
		}
	}
}

func (v *vm) cleanup() {

}

func (v *vm) disassemble() {
	addr := v.ip
	op := v.code[addr]
	name := InstructionTable[op].Name
	nargs := InstructionTable[op].Nargs

	switch nargs {
	case 0: fmt.Printf("0x%04X.%02X | %13s | \n", addr, op, name)
	case 1: fmt.Printf("0x%04X.%02X | %13s | %-8v\n", addr, op, name, v.code[v.ip + 1])
	case 2: fmt.Printf("0x%04X.%02X | %13s | %-8v %-8v\n", addr, op, name, v.code[v.ip + 1], v.code[v.ip + 2])
	case 3: fmt.Printf("0x%04X.%02X | %13s | %-8v %-8v %-8v\n", addr, op, name, v.code[v.ip + 1], v.code[v.ip + 2], v.code[v.ip + 3])
	}
}