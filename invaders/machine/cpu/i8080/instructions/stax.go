package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func STAX(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	rp := (opcode & 0x03) >> 4
	switch rp {
	case BC:
		address := uint16(registers.B)<<8 | uint16(registers.C)
		memory.Write(address, registers.A)
	case DE:
		address := uint16(registers.D)<<8 | uint16(registers.E)
		memory.Write(address, registers.A)
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
