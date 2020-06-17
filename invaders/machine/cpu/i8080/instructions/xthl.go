package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func XTHL(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	l := memory.Read(registers.SP)
	h := memory.Read(registers.SP + 1)

	memory.Write(registers.SP, registers.L)
	memory.Write(registers.SP+1, registers.H)

	registers.H = h
	registers.L = l

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
