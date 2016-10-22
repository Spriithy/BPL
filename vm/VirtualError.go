package vm

type VirtualError int

const (
	ERROR_CLEAR VirtualError = iota
	UNEXPECTED_OPCODE_ERROR
	INVALID_VIRTUAL_REFERENCE_PROCESS_ERROR
	INTEGER_VALUE_EXPECTED_ERROR
	UNSIGNED_INTEGER_EXPECTED_ERROR
	ZERO_DIVISION_ERROR
)