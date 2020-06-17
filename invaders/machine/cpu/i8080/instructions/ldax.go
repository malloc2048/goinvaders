package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func LDAX(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		registers.A = memory.Read(RegisterPairValue(BC, registers))
	case DE:
		registers.A = memory.Read(RegisterPairValue(DE, registers))
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
