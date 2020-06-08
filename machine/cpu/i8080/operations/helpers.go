package operations

func ReadMemoryAtHL() uint8 {
	address := registers[H]<<8 | registers[L]
	return memory[address]
}

func WriteMemoryAtHL(data uint8) {
	address := registers[H]<<8 | registers[L]
	memory[address] = data
}

func ReadNextByte() uint8 {
	value := memory[pc]
	pc += 1
	return value
}

func ReadNextWord() uint16 {
	return uint16(ReadNextByte()) | uint16(ReadNextByte())<<8
}

func GetRegisterPair(rp uint8) uint16 {
	switch rp {
	case BC:
		return registers[B]<<8 | registers[C]
	case DE:
		return registers[D]<<8 | registers[E]
	case HL:
		return registers[H]<<8 | registers[L]
	case SP:
		return sp
	default:
		return 0
	}
}

func SetRegisterPair(rp uint8, value uint16) {
	switch rp {
	case BC:
		registers[B] = (value & 0xff00) >> 8
		registers[C] = value & 0x00ff
	case DE:
		registers[D] = (value & 0xff00) >> 8
		registers[E] = value & 0x00ff
	case HL:
		registers[H] = (value & 0xff00) >> 8
		registers[L] = value & 0x00ff
	case SP:
		sp = value
	default:

	}
}
