package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

// RP_HIGH <- byte 3 (PC + 1), RP_LOW <- byte 2 (PC)
func LXI(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	registerPair := (opcode & 0x30) >> 4
	switch registerPair {
	case BC:
		registers.C = memory.Read(registers.PC)
		registers.B = memory.Read(registers.PC + 1)
	case DE:
		registers.E = memory.Read(registers.PC)
		registers.D = memory.Read(registers.PC + 1)
	case HL:
		registers.L = memory.Read(registers.PC)
		registers.H = memory.Read(registers.PC + 1)
	case SP:
		registers.SP = memory.ReadWord(registers.PC)
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
