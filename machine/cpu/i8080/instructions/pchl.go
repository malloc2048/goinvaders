package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
)

func PCHL(registers *registers.Registers) {
	registers.PC = uint16(registers.H)<<8 | uint16(registers.L)
}
