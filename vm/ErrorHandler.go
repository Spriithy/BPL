package vm

import "os"

type ErrorHandler struct {
	boundto     *vm
	stresscount int
}

func NewVirtualErrorHandler(bindto *vm) *ErrorHandler {
	return &ErrorHandler{bindto, 0}
}

func (e *ErrorHandler) stress(code VirtualError, msg string) {
	fatal := false

	switch code {
	case ERROR_CLEAR: return
	case INTEGER_VALUE_EXPECTED_ERROR:
		println("INTEGER_VALUE_EXPECTED_ERROR (", e.stresscount, ") report:")
		println(msg)
		fatal = true
	case UNSIGNED_INTEGER_EXPECTED_ERROR:
		println("UNSIGNED_INTEGER_VALUE_EXPECTED_ERROR (", e.stresscount, ") report:")
		println(msg)
		fatal = true
	case ZERO_DIVISION_ERROR:
		println("ZERO_DIVISION_ERROR (", e.stresscount, ") report:")
		print(msg)
		fatal = true
	case UNEXPECTED_OPCODE_ERROR:
		println("UNEXPECTED_OPCODE_ERROR (", e.stresscount, ") report:")
		println(msg)
		fatal = true
	case INVALID_VIRTUAL_REFERENCE_PROCESS_ERROR:
		println("INVALID_VIRTUAL_REFERENCE_PROCESS_ERROR (", e.stresscount, ") report:")
		println(msg)
		fatal = true
	default:
		println("UNKNOWN_ERROR (", e.stresscount, ") report:")
		println(msg)
	}
	println("    IP    |  INSTRUCTION  | ARGS")
	e.boundto.ip--
	e.boundto.disassemble()
	println("Stack:")
	println("\t", e.boundto.String())

	e.stresscount++
	if fatal {
		os.Exit(1)
	}
}
