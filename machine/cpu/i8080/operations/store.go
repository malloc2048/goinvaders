package operations

func StoreAccumulatorDirect() {
	memory[ReadNextWord()] = uint8(registers[0x07])
}

func StoreHLDirect() {
	address := ReadNextWord()
	memory[address] = uint8(registers[0x05])
	memory[address+1] = uint8(registers[0x04])
}

func StoreAccumulatorIndirect(rp uint8) {
	switch rp {
	case BC:
		memory[GetRegisterPair(BC)] = uint8(registers[A])
	case DE:
		memory[GetRegisterPair(DE)] = uint8(registers[A])
	case HL:
	case SP:
	default:
	}
}

func ExchangeHLWithDE() {
	tmp := GetRegisterPair(HL)
	SetRegisterPair(HL, GetRegisterPair(DE))
	SetRegisterPair(DE, tmp)
}
