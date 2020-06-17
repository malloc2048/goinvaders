package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func LHLD(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	address := memory.ReadWord(registers.PC)
	registers.H = memory.Read(address + 1)
	registers.L = memory.Read(address)

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
