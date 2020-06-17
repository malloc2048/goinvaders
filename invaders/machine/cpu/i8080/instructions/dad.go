package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

// HL = HL +RP
func DAD(opcode byte, registers *registers.Registers, flags *registers.Flags) {
	l := uint16(registers.L)
	h := uint16(registers.H)
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		l += uint16(registers.C)
		h += uint16(registers.B)
	case DE:
		l += uint16(registers.E)
		h += uint16(registers.D)
	case HL:
		l += uint16(registers.L)
		h += uint16(registers.H)
	case SP:
		l += registers.SP & 0x00ff
		h += registers.SP >> 8
	}
	if l&0x100 != 0 {
		h += 1
	}
	registers.L = uint8(l)
	registers.H = uint8(h)
	flags.Carry = h&0x100 != 0

	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
