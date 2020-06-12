package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func LHLD(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	address := memory.ReadWord(registers.PC)
	value := memory.ReadWord(address)

	registers.H = uint8(value >> 8)
	registers.L = uint8(value & 0xff)

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
