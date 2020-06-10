package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func SHLD(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	address := memory.ReadWord(registers.PC)
	memory.Write(address, registers.L)
	memory.Write(address+1, registers.H)

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
