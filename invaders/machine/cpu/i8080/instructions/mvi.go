package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func MVI(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	register := (opcode & 0x38) >> 4

	switch register {
	case B:
		registers.B = memory.Read(registers.PC)
	case C:
		registers.C = memory.Read(registers.PC)
	case D:
		registers.D = memory.Read(registers.PC)
	case E:
		registers.E = memory.Read(registers.PC)
	case H:
		registers.H = memory.Read(registers.PC)
	case L:
		registers.L = memory.Read(registers.PC)
	case M:
		address := uint16(registers.H)<<8 | (uint16(registers.L))
		memory.Write(address, memory.Read(registers.PC))
	case A:
		registers.A = memory.Read(registers.PC)
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
