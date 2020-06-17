package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func ORA(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	src := opcode & 0x07

	switch src {
	case B:
		registers.A |= registers.B
	case C:
		registers.A |= registers.C
	case D:
		registers.A |= registers.D
	case E:
		registers.A |= registers.E
	case H:
		registers.A |= registers.H
	case L:
		registers.A |= registers.L
	case A:
		registers.A |= registers.A
	case M:
		registers.A |= memory.Read(RegisterPairValue(HL, registers))
	}

	flags.Carry = false
	flags.HalfCarry = false
	flags.Sign = registers.A > 0x7f
	flags.Zero = registers.A == 0x00
	flags.Parity = CalculateParity(registers.A)

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
