package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
)

func SPHL(opcode byte, registers *registers.Registers) {
	registers.SP = uint16(registers.H)<<8 | uint16(registers.L)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
