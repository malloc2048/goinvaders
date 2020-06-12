package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func LDAX(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	registerPair := (opcode & 0x30) >> 4
	var address uint16

	switch registerPair {
	case BC:
		address = uint16(registers.B)<<8 | uint16(registers.C)
	case DE:
		address = uint16(registers.D)<<8 | uint16(registers.E)
	}
	registers.A = memory.Read(address)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
