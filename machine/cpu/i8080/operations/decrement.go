package operations

func Decrement(dst uint8) {
	switch dst {
	case B:
		registers[B] -= 1
	case C:
		registers[C] -= 1
	case D:
		registers[D] -= 1
	case E:
		registers[E] -= 1
	case H:
		registers[H] -= 1
	case L:
		registers[L] -= 1
	case M:
		WriteMemoryAtHL(ReadMemoryAtHL() - 1)
	case A:
		registers[A] -= 1
	}
}
