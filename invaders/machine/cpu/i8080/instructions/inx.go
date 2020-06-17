package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

// RP <- RP + 1
func INX(opcode byte, registers *registers.Registers) {
	registerPair := (opcode & 0x30) >> 4
	switch registerPair {
	case BC:
		registers.C += 1
		if registers.C == 0x00 {
			registers.B += 1
		}
	case DE:
		registers.E += 1
		if registers.E == 0x00 {
			registers.D += 1
		}
	case HL:
		registers.L += 1
		if registers.L == 0x00 {
			registers.H += 1
		}
	case SP:
		registers.SP += 1
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
