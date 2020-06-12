package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func CMP(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	var result uint16
	src := opcode & 0x07

	switch src {
	case B:
		result = uint16(registers.A) - uint16(registers.B)
	case C:
		result = uint16(registers.A) - uint16(registers.C)
	case D:
		result = uint16(registers.A) - uint16(registers.D)
	case E:
		result = uint16(registers.A) - uint16(registers.E)
	case H:
		result = uint16(registers.A) - uint16(registers.H)
	case L:
		result = uint16(registers.A) - uint16(registers.L)
	case A:
		result = uint16(registers.A) - uint16(registers.A)
	case M:
		result = uint16(registers.A) - uint16(memory.Read(RegisterPairValue(HL, registers)))
	}

	flags.Sign = result > 0x7f
	flags.Zero = result == 0x00
	flags.Carry = result&0x0100 != 0
	flags.Parity = CalculateParity(uint8(result))
	flags.HalfCarry = uint8(result)&0x0f > registers.A&0x0f

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
