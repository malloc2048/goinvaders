package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func POP(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		registers.C = memory.Read(registers.SP)
		registers.B = memory.Read(registers.SP + 1)
	case DE:
		registers.E = memory.Read(registers.SP)
		registers.D = memory.Read(registers.SP + 1)
	case HL:
		registers.L = memory.Read(registers.SP)
		registers.H = memory.Read(registers.SP + 1)
	case SP:
		restorePSW(memory.Read(registers.SP), flags)
		registers.A = memory.Read(registers.SP + 1)
	}
	registers.SP += 2
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}

func restorePSW(psw uint8, flags *registers.Flags) {
	flags.Sign = psw&0x02 != 0
	flags.Zero = psw&0x01 != 0
	flags.Carry = psw&0x08 != 0
	flags.Parity = psw&0x04 != 0
	flags.HalfCarry = psw&0x10 != 0
}

func PUSH(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		memory.Write(registers.SP-1, registers.B)
		memory.Write(registers.SP-2, registers.C)
	case DE:
		memory.Write(registers.SP-1, registers.D)
		memory.Write(registers.SP-2, registers.E)
	case HL:
		memory.Write(registers.SP-1, registers.H)
		memory.Write(registers.SP-2, registers.L)
	case SP:
		memory.Write(registers.SP-1, registers.A)
		memory.Write(registers.SP-2, makePSW(flags))
	}
	registers.SP -= 2
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}

func makePSW(flags *registers.Flags) uint8 {
	psw := uint8(0)

	if flags.Zero {
		psw |= 0x01
	}
	if flags.Sign {
		psw |= 0x02
	}
	if flags.Carry {
		psw |= 0x08
	}
	if flags.Parity {
		psw |= 0x04
	}
	if flags.HalfCarry {
		psw |= 0x10
	}

	return psw
}
