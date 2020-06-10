package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func SUB(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	var diff uint16
	src := opcode & 0x07

	switch src {
	case B:
		diff = uint16(registers.A) - uint16(registers.B)
	case C:
		diff = uint16(registers.A) - uint16(registers.C)
	case D:
		diff = uint16(registers.A) - uint16(registers.D)
	case E:
		diff = uint16(registers.A) - uint16(registers.E)
	case H:
		diff = uint16(registers.A) - uint16(registers.H)
	case L:
		diff = uint16(registers.A) - uint16(registers.L)
	case A:
		diff = uint16(registers.A) - uint16(registers.A)
	case M:
		diff = uint16(registers.A) - uint16(memory.ReadByte(RegisterPairValue(HL, registers)))
	}
	flags.Sign = diff > 0x7f
	flags.Zero = diff == 0x00
	flags.Carry = (diff & 0x100) != 0
	flags.Parity = CalculateParity(uint8(diff & 0xff))
	flags.HalfCarry = (registers.A & 0x0f) < uint8(diff&0x000f)

	registers.A = uint8(diff & 0xff)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
