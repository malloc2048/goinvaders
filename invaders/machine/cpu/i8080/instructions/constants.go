package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
)

const (
	BC = 0x00
	DE = 0x01
	HL = 0x02
	SP = 0x03
)

const (
	B = 0x00
	C = 0x01
	D = 0x02
	E = 0x03
	H = 0x04
	L = 0x05
	M = 0x06
	A = 0x07
)

const (
	NotZero    = 0x00
	Zero       = 0x01
	NoCarry    = 0x02
	Carry      = 0x03
	ParityOdd  = 0x04
	ParityEven = 0x05
	Plus       = 0x06
	Minus      = 0x07
)

var OpcodesLength = [...]uint8{
	//  0  1  2  3  4  5  6  7  8  9  A  B  C  D  E  F
	1, 3, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 0
	1, 3, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 1
	1, 3, 3, 1, 1, 1, 2, 1, 1, 1, 3, 1, 1, 1, 2, 1, // 2
	1, 3, 3, 1, 1, 1, 2, 1, 1, 1, 3, 1, 1, 1, 2, 1, // 3
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 4
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 5
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 6
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 7
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 8
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // 9
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // A
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, // B
	1, 1, 3, 3, 3, 1, 2, 1, 1, 1, 3, 1, 3, 3, 2, 1, // C
	1, 1, 3, 2, 3, 1, 2, 1, 1, 1, 3, 2, 3, 1, 2, 1, // D
	1, 1, 3, 1, 3, 1, 2, 1, 1, 1, 3, 1, 3, 1, 2, 1, // E
	1, 1, 3, 1, 3, 1, 2, 1, 1, 1, 3, 1, 3, 1, 2, 1, // F
}

func CalculateParity(value byte) bool {
	bitsSet := 0
	for i := 0; i < 8; i++ {
		if (value>>i)&0x01 == 1 {
			bitsSet++
		}
	}
	return bitsSet&0x01 == 0x00
}

func RegisterPairValue(src uint8, registers *registers.Registers) uint16 {
	switch src {
	case BC:
		return uint16(registers.B)<<8 | uint16(registers.C)
	case DE:
		return uint16(registers.D)<<8 | uint16(registers.E)
	case HL:
		//log.Printf("%04x\n", uint16(registers.H)<<8 | uint16(registers.L))
		return uint16(registers.H)<<8 | uint16(registers.L)
	default:
		return 0
	}
}
