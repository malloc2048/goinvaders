package operations

func LoadRegisterPairImmediate(dst uint8) {
	switch dst {
	case BC:
		registers[C] = uint16(ReadNextByte())
		registers[B] = uint16(ReadNextByte())
	case DE:
		registers[E] = uint16(ReadNextByte())
		registers[D] = uint16(ReadNextByte())
	case HL:
		registers[L] = uint16(ReadNextByte())
		registers[H] = uint16(ReadNextByte())
	case SP:
		sp = uint16(uint16(ReadNextByte()) | uint16(ReadNextByte())<<8)
	default:
		// TODO: maybe add some debug log here, just in case
	}
}

func LoadAccumulator() {
	registers[A] = uint16(memory[ReadNextWord()])
}

func LoadHLDirect() {
	address := ReadNextWord()
	registers[L] = uint16(memory[address])
	registers[H] = uint16(memory[address+1])
}

func LoadAccumulatorIndirect(rp uint8) {
	switch rp {
	case BC:
		registers[A] = uint16(memory[GetRegisterPair(BC)])
	case DE:
		registers[A] = uint16(memory[GetRegisterPair(DE)])
	case HL:
	case SP:
	default:
	}
}
