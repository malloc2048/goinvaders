package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

func PCHL(registers *registers.Registers) {
	registers.PC = RegisterPairValue(HL, registers)
}
