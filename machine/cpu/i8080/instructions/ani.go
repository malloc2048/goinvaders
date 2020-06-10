package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func ANI(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	registers.A = registers.A & memory.ReadByte(registers.PC)

	flags.Carry = false
	flags.HalfCarry = false
	flags.Sign = registers.A > 0x7f
	flags.Zero = registers.A == 0x00
	flags.Parity = CalculateParity(registers.A)

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
