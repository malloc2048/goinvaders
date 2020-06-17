package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

func DCX(opcode byte, registers *registers.Registers) {
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		if registers.C == 0x00 {
			registers.B -= 1
		}
		registers.C -= 1
	case DE:
		if registers.E == 0x00 {
			registers.D -= 1
		}
		registers.E -= 1
	case HL:
		if registers.L == 0x00 {
			registers.H -= 1
		}
		registers.L -= 1
	case SP:
		registers.SP -= 1
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
