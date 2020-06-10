package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func XRA(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	var result uint8
	src := opcode & 0x07

	switch src {
	case B:
		result = registers.B ^ registers.A
	case C:
		result = registers.C ^ registers.A
	case D:
		result = registers.D ^ registers.A
	case E:
		result = registers.E ^ registers.A
	case H:
		result = registers.H ^ registers.A
	case L:
		result = registers.L ^ registers.A
	case A:
		result = registers.A ^ registers.A
	case M:
		result = registers.A ^ memory.ReadByte(RegisterPairValue(HL, registers))
	}

	flags.Carry = false
	flags.HalfCarry = false
	flags.Sign = result > 0x7f
	flags.Zero = result == 0x00
	flags.Parity = CalculateParity(result)

	registers.A = result
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
