package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

// (RP) <- A
func STAX(opcode byte, memory *memory.Memory, registers *registers.Registers) {
	rp := (opcode & 0x03) >> 4
	switch rp {
	case BC:
		memory.Write(RegisterPairValue(BC, registers), registers.A)
	case DE:
		memory.Write(RegisterPairValue(DE, registers), registers.A)
	}
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
