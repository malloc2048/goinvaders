package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

func Rotate(opcode byte, registers *registers.Registers, flags *registers.Flags) {
	switch opcode {
	case 0x07: // RLC
		carry := registers.A >> 7
		registers.A = registers.A<<0x01 | carry
		flags.Carry = carry == 0x01
	case 0x0f:
		carry := registers.A & 0x01
		registers.A = registers.A>>0x01 | carry<<0x07
		flags.Carry = carry == 0x01
	case 0x17:
		carry := flags.Carry
		flags.Carry = registers.A>>0x07 == 0x01
		if carry {
			registers.A = registers.A<<0x01 | 0x01
		} else {
			registers.A = registers.A << 0x01
		}
	case 0x1f:
		carry := flags.Carry
		flags.Carry = registers.A&0x01 == 0x01
		if carry {
			registers.A = registers.A>>0x01 | 0x80
		} else {
			registers.A = registers.A >> 0x01
		}
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
