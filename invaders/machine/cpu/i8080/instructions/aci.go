package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func ACI(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	sum := uint16(registers.A) + uint16(memory.Read(registers.PC))

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
