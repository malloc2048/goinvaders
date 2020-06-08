package operations

func reset() {
	for i := uint32(0); i < 0x10000; i++ {
		memory[i] = 0x00
	}

	for i := uint8(0); i < 0x08; i++ {
		registers[i] = 0x0000
	}

	pc = 0x0000
	sp = 0x0000
	SetSign(false)
	SetZero(false)
	SetCarry(false)
	SetParity(false)
	SetAuxiliary(false)
}
