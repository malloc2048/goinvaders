package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
)

func DAA(opcode byte, registers *registers.Registers, flags *registers.Flags) {
	lsb := registers.A & 0x0f
	if flags.HalfCarry || lsb > 9 {
		registers.A += 6
	}

	msb := registers.A >> 4
	if flags.Carry || msb > 9 {
		registers.A += 0x60
	}

	flags.Carry = msb > 9
	flags.Zero = registers.A == 0
	flags.Sign = registers.A > 0x7f
	flags.HalfCarry = (lsb & 0x10) != 0
	flags.Parity = CalculateParity(registers.A)

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
