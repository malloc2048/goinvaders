package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

func XCHG(opcode byte, registers *registers.Registers) {
	d := registers.D
	registers.D = registers.H
	registers.H = d

	e := registers.E
	registers.E = registers.L
	registers.L = e

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
