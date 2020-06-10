package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func LXI(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	registerPair := (opcode & 0x30) >> 4
	switch registerPair {
	case BC:
		registers.C = memory.ReadByte(registers.PC)
		registers.B = memory.ReadByte(registers.PC + 1)
	case DE:
		registers.E = memory.ReadByte(registers.PC)
		registers.D = memory.ReadByte(registers.PC + 1)
	case HL:
		registers.L = memory.ReadByte(registers.PC)
		registers.H = memory.ReadByte(registers.PC + 1)
	case SP:
		registers.SP = memory.ReadWord(registers.PC)
		//registers.SP = uint16(memory.ReadByte(registers.PC)) << 8 | uint16(memory.ReadByte(registers.PC + 1))
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
