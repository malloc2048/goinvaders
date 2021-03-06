package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

// R <- R - 1
func DCR(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	register := (opcode & 0x38) >> 3
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
		address := RegisterPairValue(HL, registers)
		value = memory.Read(address) - 1
		memory.Write(address, value)
	case A:
		registers.A -= 1
		value = registers.A
	}
	flags.Sign = value > 0x7f
	flags.Zero = value == 0x00
	flags.Parity = CalculateParity(value)
	flags.HalfCarry = (value & 0x0f) == 0x0f

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
