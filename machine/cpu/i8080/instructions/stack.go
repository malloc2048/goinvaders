package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func POP(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	registerPair := (opcode & 0x30) >> 4

	switch registerPair {
	case BC:
		registers.B = memory.ReadByte(registers.SP)
		registers.C = memory.ReadByte(registers.SP + 1)
	case DE:
		registers.D = memory.ReadByte(registers.SP)
		registers.E = memory.ReadByte(registers.SP + 1)
	case HL:
		registers.H = memory.ReadByte(registers.SP)
		registers.L = memory.ReadByte(registers.SP + 1)
	case SP:
		restorePSW(memory.ReadByte(registers.SP), flags)
		registers.A = memory.ReadByte(registers.SP + 1)
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
