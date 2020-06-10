package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
)

func DCX(opcode byte, registers *registers.Registers) {
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		value := (uint16(registers.B)<<8 | uint16(registers.C)) - 1
		registers.B = uint8(value >> 8)
		registers.C = uint8(value & 0xff)
	case DE:
		value := (uint16(registers.D)<<8 | uint16(registers.E)) - 1
		registers.D = uint8(value >> 8)
		registers.E = uint8(value & 0xff)
	case HL:
		value := (uint16(registers.H)<<8 | uint16(registers.L)) - 1
		registers.H = uint8(value >> 8)
		registers.L = uint8(value & 0xff)
	case SP:
		registers.SP -= 1
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
