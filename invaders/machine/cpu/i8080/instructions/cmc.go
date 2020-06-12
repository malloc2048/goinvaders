package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

func CMC(opcode byte, registers *registers.Registers, flags *registers.Flags) {
	flags.Carry = !flags.Carry
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
