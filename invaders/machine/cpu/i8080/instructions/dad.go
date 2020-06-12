package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

func DAD(opcode byte, registers *registers.Registers, flags *registers.Flags) {
	var sum uint32
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		sum = uint32(registers.H)<<8 | uint32(registers.L) + uint32(registers.B)<<8 | uint32(registers.C)
	case DE:
		sum = uint32(registers.H)<<8 | uint32(registers.L) + uint32(registers.D)<<8 | uint32(registers.E)
	case HL:
		sum = uint32(registers.H)<<8 | uint32(registers.L) + uint32(registers.H)<<8 | uint32(registers.L)
	case SP:
		sum = uint32(registers.H)<<8 | uint32(registers.L) + uint32(registers.SP)
	}
	registers.H = uint8(sum>>8) & 0xff
	registers.L = uint8(sum) & 0xff
	flags.Carry = (sum & 0x10000) == 1

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
