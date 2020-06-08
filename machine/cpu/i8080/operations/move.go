package operations

func MoveRegister(src, dst uint8) {
	if src != dst {
		if src == 0x06 {
			registers[dst] = uint16(ReadMemoryAtHL())
		} else if dst == 0x06 {
			WriteMemoryAtHL(uint8(registers[src] & 0xff))
		} else {
			registers[dst] = registers[src]
		}
	}
}

func MoveImmediate(dst uint8) {
	if dst == 0x06 {
		WriteMemoryAtHL(ReadNextByte())
	} else {
		registers[dst] = uint16(ReadNextByte())
	}
}
