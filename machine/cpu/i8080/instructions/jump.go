package instructions

import (
	"goinvaders/machine/cpu/i8080/registers"
	"goinvaders/machine/memory"
)

func Jump(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	var takeJump = true

	if opcode&0x07 == 0 {
		condition := opcode & 0x38 >> 3
		switch condition {
		case NotZero:
			takeJump = !flags.Zero
		case Zero:
			takeJump = flags.Zero
		case NoCarry:
			takeJump = !flags.Carry
		case Carry:
			takeJump = flags.Carry
		case ParityOdd:
			takeJump = !flags.Parity
		case ParityEven:
			takeJump = flags.Parity
		case Plus:
			takeJump = !flags.Sign
		case Minus:
			takeJump = flags.Sign
		}
	}

	if takeJump {
		registers.PC = memory.ReadWord(registers.PC)
	} else {
		registers.PC += uint16(OpcodesLength[opcode] - 1)
	}
}
