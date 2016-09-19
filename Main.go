package main

import (
	"./vm"
)

func main() {
	prog := []int {
		vm.ALC, 2,
		vm.PSH, 10,
		vm.LST, 0,
		vm.LGT, 0,
		vm.ARG, 0,
		vm.HLT,
	}

	m := vm.InitMachine(prog, vm.NativeInt(-4))
	m.Print()
	m.Start()
	m.Print()
}
