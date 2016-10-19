package compiler

type SymTable map[Hash]SymVal

type SymUsage int

const (
	UNKNOWN_USAGE SymUsage = iota
	KEYWORD_USAGE
	VARIABLE_USAGE
	FUNCTION_USAGE
	STRUCT_USAGE
)

type VarType int

const (
	UNKNOWN_VARTYP VarType = iota
	VOID_VARTYPE
	BOOL_VARTYP
	BYTE_VARTYP
	UINT_VARTYP
	INT_VARTYP
	ULONG_VARTYP
	LONG_VARTYP
	FLOAT_VARTYP
	DOUBLE_VARTYP
	COMPLEX_VARTYP
	OBJECT_VARTYP
)

type Hash uint64
type Symbol string

type SymVal struct {
	Tok     *Token
	Usage   SymUsage

	// If the symbol is a variable these only are relevant fields
	init    bool

	typ     VarType

	// If the symbol is a function only these are relevant fields
	public  bool
	ret     VarType
	nargs   int
	argst   []VarType
	syms    SymTable

	// For structures
	fields  SymTable
	methods SymTable
	statics SymTable
	natives SymTable
}

func HashSymbol(key Symbol) Hash {
	var hash Hash = 5381

	for c := range key {
		hash = (Hash(hash << 5) + hash) + Hash(c)
	}

	return hash
}

