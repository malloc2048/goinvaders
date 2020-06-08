package operations

func Subtract(src uint8) {
	switch src {
	case B:
		registers[A] -= registers[B]
	case C:
		registers[A] -= registers[C]
	case D:
		registers[A] -= registers[D]
	case E:
		registers[A] -= registers[E]
	case H:
		registers[A] -= registers[H]
	case L:
		registers[A] -= registers[L]
	case M:
		registers[A] -= uint16(ReadMemoryAtHL())
	case A:
		registers[A] -= registers[A]
	default:
	}
	SetFlags(registers[A])
	registers[A] &= 0x00ff
}

func SubtractImmediate() {
	registers[A] -= uint16(ReadNextByte())
	SetFlags(registers[A])
}

func SubtractWithBorrow(src uint8) {
	carry := Carry()
	Add(src)
	if carry {
		registers[A] -= 1
	}
}

func SubtractImmediateWithBorrow() {
	registers[A] -= uint16(ReadNextByte())
	if Carry() {
		registers[A] -= 1
	}
	SetFlags(registers[A])
}
