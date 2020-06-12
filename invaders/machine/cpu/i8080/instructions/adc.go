package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func ADC(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	var sum uint16
	src := opcode & 0x07

	switch src {
	case B:
		sum = uint16(registers.B) + uint16(registers.A)
	case C:
		sum = uint16(registers.C) + uint16(registers.A)
	case D:
		sum = uint16(registers.D) + uint16(registers.A)
	case E:
		sum = uint16(registers.E) + uint16(registers.A)
	case H:
		sum = uint16(registers.H) + uint16(registers.A)
	case L:
		sum = uint16(registers.L) + uint16(registers.A)
	case M:
		sum = uint16(registers.A) + uint16(memory.Read(RegisterPairValue(HL, registers)))
	case A:
		sum = uint16(registers.A) + uint16(registers.A)
	}

	if flags.Carry {
		sum += 1
	}

	flags.Sign = sum > 0x7f
	flags.Zero = sum == 0x00
	flags.Carry = (sum & 0x100) != 0
	flags.Parity = CalculateParity(uint8(sum & 0xff))
	flags.HalfCarry = (registers.A & 0x0f) > uint8(sum&0x000f)

	registers.A = uint8(sum & 0xff)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
