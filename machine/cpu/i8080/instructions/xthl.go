package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func XTHL(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	sp := memory.ReadWord(registers.SP)

	memory.Write(registers.SP, registers.L)
	memory.Write(registers.SP+1, registers.H)

	registers.H = uint8(sp >> 8)
	registers.L = uint8(sp & 0xff)

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
