package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func LDA(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	address := memory.ReadWord(registers.PC)
	registers.A = memory.ReadByte(address)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
