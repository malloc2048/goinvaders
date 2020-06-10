package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func STA(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	address := memory.ReadWord(registers.PC)
	memory.Write(address, registers.A)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
