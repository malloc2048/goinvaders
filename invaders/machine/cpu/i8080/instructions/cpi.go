package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func CPI(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	diff := uint16(registers.A) - uint16(memory.Read(registers.PC))
	flags.Sign = diff > 0x7f
	flags.Zero = diff == 0x00
	flags.Carry = diff&0x0100 != 0
	flags.Parity = CalculateParity(uint8(diff))
	flags.HalfCarry = uint8(diff)&0x0f > registers.A&0x0f

	//registers.PC = uint16(registers.H)<<8 | uint16(registers.L)
	registers.PC += uint16(OpcodesLength[opcode] - 1)

}
