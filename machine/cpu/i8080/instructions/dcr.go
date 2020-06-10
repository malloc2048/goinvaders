package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func DCR(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	register := (opcode & 0x30) >> 4
	var value uint8

	switch register {
	case B:
		registers.B -= 1
		value = registers.B
	case C:
		registers.C -= 1
		value = registers.C
	case D:
		registers.D -= 1
		value = registers.D
	case E:
		registers.E -= 1
		value = registers.E
	case H:
		registers.H -= 1
		value = registers.H
	case L:
		registers.L -= 1
		value = registers.L
	case M:
		address := uint16(registers.H)<<8 | (uint16(registers.L))
		value = memory.ReadByte(address) - 1
		memory.Write(address, memory.ReadByte(address)-1)
	case A:
		registers.A -= 1
		value = registers.A
	}
	flags.Sign = value > 0x7f
	flags.Zero = value == 0x00
	flags.Parity = CalculateParity(value)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
