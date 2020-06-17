package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func STA(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	memory.Write(memory.ReadWord(registers.PC), registers.A)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
