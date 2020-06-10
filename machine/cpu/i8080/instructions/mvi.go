package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func MVI(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	register := (opcode & 0x30) >> 4

	switch register {
	case B:
		registers.B = memory.ReadByte(registers.PC)
	case C:
		registers.C = memory.ReadByte(registers.PC)
	case D:
		registers.D = memory.ReadByte(registers.PC)
	case E:
		registers.E = memory.ReadByte(registers.PC)
	case H:
		registers.H = memory.ReadByte(registers.PC)
	case L:
		registers.L = memory.ReadByte(registers.PC)
	case M:
		address := uint16(registers.H)<<8 | (uint16(registers.L))
		memory.Write(address, memory.ReadByte(registers.PC))
	case A:
		registers.A = memory.ReadByte(registers.PC)
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
