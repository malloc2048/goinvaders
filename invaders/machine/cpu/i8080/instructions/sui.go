package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func SUI(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	diff := uint16(registers.A) - uint16(memory.Read(registers.PC))

	flags.Sign = diff > 0x7f
	flags.Zero = diff == 0x00
	flags.Carry = (diff & 0x100) != 0
	flags.Parity = CalculateParity(uint8(diff & 0xff))
	flags.HalfCarry = (registers.A & 0x0f) < uint8(diff&0x000f)

	registers.A = uint8(diff & 0xff)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
