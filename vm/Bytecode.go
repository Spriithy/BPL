package vm

type Bytecode int

const (
	NULLREF Bytecode = iota
	HALT

	LNFEED
	STDIN // Reference to string containing input
	STDOUT_VAL
	STDOUT_REF
	STDOUT_UNI
	LNOUT_VAL
	LNOUT_REF
	LNOUT_UNI

	FOPEN
	FCLOSE

	CLOCK // Gives time millis (Real)
	RRAND // Ranges 0-1
	IRAND // ARG0, ARG1 = [INF, SUP[

	SIGOF
	IABS
	RABS

	CEIL
	ROUND
	FLOOR

	IMAX_N // All those 4 need <N> argument
	RMAX_N
	IMIN_N
	RMIN_N

	SIN
	COS
	TAN
	ASIN
	ACOS
	ATAN
	ATAN2

	IPOW
	SQRT
	POW

	EXP
	LOG_2
	LOG_E
	LOG_10
	LOG_X

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

	RCONST_0
	RCONST_1
	RCONST_E
	RCONST_PI
	RCONST_PHI
	RCONST_N
	RSTORE_0
	RSTORE_1
	RSTORE_2
	RSTORE_3
	RSTORE_N
	RLOAD_0
	RLOAD_1
	RLOAD_2
	RLOAD_3
	RLOAD_N

	EQ
	NEQ
	CMP
	LT
	LEQ
	GT
	GEQ
	NEG
	INC
	DEC
	DIV

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

	RADD
	RSUB
	RMUL
	RSHR
	RSHL

	// Technical Opcodes
	NOP
	DROP
	DUP
	DUP2
	SWAP
	SWAP2

	// ARG0 = Next IP if validated
	BR
	BR_0
	BR_N0
	BR_LT
	BR_GT
	BR_LEQ
	BR_GEQ
	BR_EQ
	BR_NEQ
	BR_NUL
	BR_NNUL

	// ARG0 = Next IP if validate, ARG1 if not
	IF_0
	IF_N0
	IF_LT
	IF_GT
	IF_LEQ
	IF_GEQ
	IF_EQ
	IF_NEQ
	IF_NUL
	IF_NNUL

	CALL // NARGS
	RETURN // no value
	RET
	RET_0
	RET_N0
	RET_NUL
	RET_NNUL
)

type Instruction struct {
	Name  string
	Nargs int
}

var InstructionTable = map[Bytecode]Instruction{
	NULLREF: {"null_ref", 0},
	HALT: {"halt", 0},

	STDIN: {"stdin", 1},
	LNFEED: {"lnfeed", 0},
	STDOUT_VAL: {"stdout", 0},
	STDOUT_REF: {"stdout_rf", 0},
	STDOUT_UNI: {"stdout", 0},
	LNOUT_VAL: {"lnout_val", 0},
	LNOUT_REF: {"lnout_rf", 0},
	LNOUT_UNI: {"lnout_uni", 0},

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

	INC: {"inc", 1}, DEC: {"dec", 1},
	DIV: {"div", 0}, CMP: {"icmp", 0},
	LT: {"lt", 0}, LEQ: {"leq", 0},
	GT: {"gt", 0}, GEQ: {"geq", 0},
	EQ: {"eq", 0}, NEQ: {"neq", 0},

	NEG: {"neg", 0}, IADD: {"iadd", 0},
	ISUB: {"isub", 0}, IMUL: {"imul", 0},
	ISHR: {"ishr", 0}, ISHL: {"ishl", 0},
	IAND: {"iand", 0}, IOR: {"ior", 0},
	IXOR: {"ixor", 0}, INOT: {"inot", 0}, IMOD: {"imod", 0},
	ICOMPL1: {"icompl1", 0}, ICOMPL2: {"icompl2", 0},

	RADD: {"radd", 0}, RSUB: {"rsub", 0},
	RMUL: {"rmul", 0},
	RSHR: {"rshr", 0}, RSHL: {"rshl", 0},

	CLOCK: {"clock", 0},

	RRAND: {"rrand", 0}, IRAND: {"irand", 2}, // irand \in [Min, Max]
	SIGOF: {"signof", 0}, IABS: {"iabs", 0}, RABS: {"rabs", 0},
	CEIL: {"ceil", 0}, ROUND: {"round", 0}, FLOOR: {"floor", 0},

	// All those need <N> parameter
	IMAX_N: {"imax_n", 1}, RMAX_N: {"rmax_n", 1},
	IMIN_N: {"imin_n", 1}, RMIN_N: {"rmin_n", 1},

	SIN: {"sin", 0},
	COS: {"cos", 0}, TAN: {"tan", 0},
	ASIN: {"asin", 0}, ACOS: {"acos", 0},
	ATAN: {"atan", 0}, ATAN2: {"atan2", 0},

	SQRT: {"sqrt", 0}, IPOW: {"ipow", 0},
	POW: {"pow", 0}, EXP: {"exp", 0},
	LOG_2: {"log_2", 0}, LOG_E: {"log_E", 0},
	LOG_10: {"log_10", 0}, LOG_X: {"log_x", 0},

	BR: {"branch", 1}, BR_0: {"branch_0", 1},
	IF_0: {"if_0", 2}, IF_N0: {"if_n0", 2},
}
