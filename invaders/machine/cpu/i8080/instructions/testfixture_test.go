package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func MakeTestFixture() *TestFixture {
	return &TestFixture{
		memory: new(memory.Memory),
		flags:  new(registers.Flags),
		regs:   new(registers.Registers),
	}
}

type TestFixture struct {
	memory *memory.Memory
	flags  *registers.Flags
	regs   *registers.Registers
}
