package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func RST(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	rstVector := (opcode & 0x38) >> 3

	memory.Write(registers.SP-1, uint8(registers.PC>>8))
	memory.Write(registers.SP-2, uint8(registers.PC&0xff))
	registers.SP -= 2

	registers.PC = uint16(rstVector * 8)
}
