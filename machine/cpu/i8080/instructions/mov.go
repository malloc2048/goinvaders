package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func MOV(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	src := opcode & 0x07
	dst := (opcode & 0x38) >> 3

	switch src {
	case B:
		move(dst, registers.B, memory, registers)
	case C:
		move(dst, registers.C, memory, registers)
	case D:
		move(dst, registers.D, memory, registers)
	case E:
		move(dst, registers.E, memory, registers)
	case H:
		move(dst, registers.H, memory, registers)
	case L:
		move(dst, registers.L, memory, registers)
	case M:
		address := uint16(registers.H)<<8 | (uint16(registers.L))
		move(dst, memory.ReadByte(address), memory, registers)
	case A:
		move(dst, registers.B, memory, registers)
	}

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}

func move(dst uint8, value uint8, memory *memory.Memory, registers *registers.Registers) {
	switch dst {
	case B:
		registers.B = value
	case C:
		registers.C = value
	case D:
		registers.D = value
	case E:
		registers.E = value
	case H:
		registers.H = value
	case L:
		registers.L = value
	case M:
		address := uint16(registers.H)<<8 | (uint16(registers.L))
		memory.Write(address, value)
	case A:
		registers.A = value
	}
}
