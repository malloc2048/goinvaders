package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

func CMA(opcode byte, registers *registers.Registers) {
	registers.A ^= 0xff
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
