package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
)

func STC(opcode byte, registers *registers.Registers, flags *registers.Flags) {
	flags.Carry = true
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
