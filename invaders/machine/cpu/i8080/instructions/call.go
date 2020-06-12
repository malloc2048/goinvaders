package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"goinvaders/invaders/machine/memory"
)

func Call(opcode byte, memory *memory.Memory, registers *registers.Registers, flags *registers.Flags) {
	var takeJump = true

	if opcode&0x07 == 2 {
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
		address := registers.PC + 2
		memory.Write(registers.SP-1, uint8(address>>8))
		memory.Write(registers.SP-2, uint8(address&0xff))

		registers.PC = memory.ReadWord(registers.PC)
		registers.SP -= 2
	} else {
		registers.PC += uint16(OpcodesLength[opcode] - 1)
	}
}
