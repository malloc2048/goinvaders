package operations

func SetFlags(value uint16) {
	SetZero(value == 0)
	SetSign(value > 0x0077)
	SetParity(parity(value & 0x00ff))
	SetCarry(value > 0x00ff)

	// TODO: half carry
}

func parity(value uint16) bool {
	setBits := 0
	for i := 0; i < 8; i++ {
		if value>>i&1 == 1 {
			setBits += 1
		}
	}
	return setBits&0x0001 != 1
}

func Zero() bool {
	return (flags & 0x01) > 0
}

func Sign() bool {
	return (flags & 0x02) > 0
}

func Parity() bool {
	return (flags & 0x04) > 0
}

func Carry() bool {
	return (flags & 0x08) > 0
}

func Auxiliary() bool {
	return (flags & 0x10) > 0
}

func SetZero(zero bool) {
	if zero {
		flags = flags | 0x01
	} else {
		flags = flags & 0xfe
	}
}

func SetSign(sign bool) {
	if sign {
		flags = flags | 0x02
	} else {
		flags = flags & 0xfd
	}
}

func SetParity(parity bool) {
	if parity {
		flags = flags | 0x04
	} else {
		flags = flags & 0xfb
	}
}

func SetCarry(carry bool) {
	if carry {
		flags = flags | 0x08
	} else {
		flags = flags & 0xf7
	}
}

func SetAuxiliary(auxiliary bool) {
	if auxiliary {
		flags = flags | 0x10
	} else {
		flags = flags & 0xef
	}
}
