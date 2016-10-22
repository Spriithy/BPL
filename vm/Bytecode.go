package vm

type Bytecode byte

const (
	NULL Bytecode = iota
	HALT

	LNFEED
	PRINT_VAL
	PRINT_REF
	PRINT_UNI
	PRINTLN_VAL
	PRINTLN_REF
	PRINTLN_UNI

	ICONST_0
	ICONST_1
	ICONST_2
	ICONST_3
	ICONST_4
	ICONST_5
	ICONST_N
	ISTORE_0
	ISTORE_1
	ISTORE_2
	ISTORE_3
	ISTORE_N
	ILOAD_0
	ILOAD_1
	ILOAD_2
	ILOAD_3
	ILOAD_N

	CMP
	LT
	LEQT
	GT
	GEQT
	NEG
	INC
	DEC
	DIV

	IEQ
	INEQ
	IADD
	ISUB
	IMUL
	IMOD
	ISHR
	ISHL
	IAND
	IOR
	IXOR
	ICOMPL1
	ICOMPL2
	INOT

	RCONST_0
	RCONST_1

	RCONST_E
	RCONST_PI
	RCONST_PHI
	RCONST_N

	REQ
	RNEQ
	RADD
	RSUB
	RMUL
	RMOD
	RSHR
	RSHL
	RAND
	ROR
	RXOR
	RCOMPL1
	RCOMPL2
	RNOT

	// RSTOR_<N> : Store real local offset <N>
	RSTORE_0
	RSTORE_1
	RSTORE_2
	RSTORE_3
	RSTORE_N

	// RLOAD_<N> : Load real local offset <N>
	RLOAD_0
	RLOAD_1
	RLOAD_2
	RLOAD_3
	RLOAD_N
)

type Instruction struct {
	Name  string
	Nargs int
}

var InstructionTable = map[Bytecode]Instruction{
	NULL: {"null", 0},
	HALT: {"halt", 0},

	LNFEED: {"lnfeed", 0},
	PRINT_VAL: {"print_val", 0},
	PRINT_REF: {"print_ref", 0},
	PRINT_UNI: {"print_uni", 0},
	PRINTLN_VAL: {"println_val", 0},
	PRINTLN_REF: {"println_ref", 0},
	PRINTLN_UNI: {"println_uni", 0},

	ICONST_0: {"iconst_0", 0}, ICONST_1: {"iconst_1", 0},
	ICONST_2: {"iconst_2", 0}, ICONST_3: {"iconst_3", 0},
	ICONST_4: {"iconst_4", 0}, ICONST_5: {"iconst_5", 0},
	ICONST_N: {"iconst", 1}, // ARG0 = IPool[ID] for constant

	ISTORE_0: {"istore_0", 0}, ISTORE_1: {"istore_1", 0},
	ISTORE_2: {"istore_2", 0}, ISTORE_3: {"istore_3", 0},
	ISTORE_N: {"istore", 1}, // ARG0 = Local variable offset

	ILOAD_0: {"iload_0", 0}, ILOAD_1: {"iload_1", 0},
	ILOAD_2: {"iload_2", 0}, ILOAD_3: {"iload_3", 0},
	ILOAD_N: {"iload", 1}, // ARG0 = Local variable offset

	RCONST_0: {"rconst_0", 0}, RCONST_1: {"rconst_1", 0},
	RCONST_N: {"rconst", 1}, // ARG0 = IPool[ID] for constant

	RSTORE_0: {"rstore_0", 0}, RSTORE_1: {"rstore_1", 0},
	RSTORE_2: {"rstore_2", 0}, RSTORE_3: {"rstore_3", 0},
	RSTORE_N: {"rstore", 1}, // ARG0 = Local variable offset

	RLOAD_0: {"rload_0", 0}, RLOAD_1: {"rload_1", 0},
	RLOAD_2: {"rload_2", 0}, RLOAD_3: {"rload_3", 0},
	RLOAD_N: {"rload", 1}, // ARG0 = Local variable offset

	CMP: {"icmp", 0}, IEQ: {"ieq", 0},
	LT: {"lt", 0}, LEQT: {"leqt", 0},
	GT: {"gt", 0}, GEQT: {"geqt", 0},
	INC: {"inc", 1}, DEC: {"dec", 1},
	NEG: {"neg", 0}, IADD: {"iadd", 0},
	ISUB: {"isub", 0}, IMUL: {"imul", 0},
	DIV: {"div", 0}, IMOD: {"imod", 0},
	ISHR: {"ishr", 0}, ISHL: {"ishl", 0},
	IAND: {"iand", 0}, IOR: {"ior", 0},
	IXOR: {"ixor", 0}, INOT: {"inot", 0},
	ICOMPL1: {"icompl1", 0}, ICOMPL2: {"icompl2", 0},
}


