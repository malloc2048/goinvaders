package instructions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// A <- A + REG + CY
func TestADD_B(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x01
		fixture.regs.B = uint8(i)
		fixture.flags.Carry = true
		ADD(0x80, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+0x01, fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		assert.False(t, fixture.flags.Zero)
		assert.Equal(t, CalculateParity(uint8(i)+0x01), CalculateParity(fixture.regs.A))

		if i < 0x7f {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}

		if i == 0xff {
			assert.True(t, fixture.flags.Carry)
		} else {
			assert.False(t, fixture.flags.Carry)
		}

		if i&0xf == 0xf {
			assert.True(t, fixture.flags.HalfCarry)
		} else {
			assert.False(t, fixture.flags.HalfCarry)
		}
	}
}

func TestADD_C(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x01
		fixture.regs.C = uint8(i)
		fixture.flags.Carry = true
		ADD(0x81, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+0x01, fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		assert.False(t, fixture.flags.Zero)
		assert.Equal(t, CalculateParity(uint8(i)+0x01), CalculateParity(fixture.regs.A))

		if i < 0x7f {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}

		if i == 0xff {
			assert.True(t, fixture.flags.Carry)
		} else {
			assert.False(t, fixture.flags.Carry)
		}

		if i&0xf == 0xf {
			assert.True(t, fixture.flags.HalfCarry)
		} else {
			assert.False(t, fixture.flags.HalfCarry)
		}
	}
}

func TestADD_D(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x01
		fixture.regs.D = uint8(i)
		fixture.flags.Carry = true
		ADD(0x82, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+0x01, fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		assert.False(t, fixture.flags.Zero)
		assert.Equal(t, CalculateParity(uint8(i)+0x01), CalculateParity(fixture.regs.A))

		if i < 0x7f {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}

		if i == 0xff {
			assert.True(t, fixture.flags.Carry)
		} else {
			assert.False(t, fixture.flags.Carry)
		}

		if i&0xf == 0xf {
			assert.True(t, fixture.flags.HalfCarry)
		} else {
			assert.False(t, fixture.flags.HalfCarry)
		}
	}
}

func TestADD_E(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x01
		fixture.regs.E = uint8(i)
		fixture.flags.Carry = true
		ADD(0x83, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+0x01, fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		assert.False(t, fixture.flags.Zero)
		assert.Equal(t, CalculateParity(uint8(i)+0x01), CalculateParity(fixture.regs.A))

		if i < 0x7f {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}

		if i == 0xff {
			assert.True(t, fixture.flags.Carry)
		} else {
			assert.False(t, fixture.flags.Carry)
		}

		if i&0xf == 0xf {
			assert.True(t, fixture.flags.HalfCarry)
		} else {
			assert.False(t, fixture.flags.HalfCarry)
		}
	}
}

func TestADD_H(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x01
		fixture.regs.H = uint8(i)
		fixture.flags.Carry = true
		ADD(0x84, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+0x01, fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		assert.False(t, fixture.flags.Zero)
		assert.Equal(t, CalculateParity(uint8(i)+0x01), CalculateParity(fixture.regs.A))

		if i < 0x7f {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}

		if i == 0xff {
			assert.True(t, fixture.flags.Carry)
		} else {
			assert.False(t, fixture.flags.Carry)
		}

		if i&0xf == 0xf {
			assert.True(t, fixture.flags.HalfCarry)
		} else {
			assert.False(t, fixture.flags.HalfCarry)
		}
	}
}

func TestADD_L(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x01
		fixture.regs.L = uint8(i)
		fixture.flags.Carry = true
		ADD(0x85, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+0x01, fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		assert.False(t, fixture.flags.Zero)
		assert.Equal(t, CalculateParity(uint8(i)+0x01), CalculateParity(fixture.regs.A))

		if i < 0x7f {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}

		if i == 0xff {
			assert.True(t, fixture.flags.Carry)
		} else {
			assert.False(t, fixture.flags.Carry)
		}

		if i&0xf == 0xf {
			assert.True(t, fixture.flags.HalfCarry)
		} else {
			assert.False(t, fixture.flags.HalfCarry)
		}
	}
}

func TestADD_M(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.flags.Carry = true
		fixture.memory.Storage[0x0000] = uint8(i)
		ADD(0x86, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i), fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		if i == 0x00 {
			assert.True(t, fixture.flags.Zero)
		} else {
			assert.False(t, fixture.flags.Zero)
		}

		assert.Equal(t, CalculateParity(uint8(i)), CalculateParity(fixture.regs.A))

		if i < 0x80 {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}
	}
}

func TestADD_A(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = uint8(i)
		fixture.flags.Carry = true
		ADD(0x87, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+uint8(i), fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)

		if i == 0x00 {
			assert.True(t, fixture.flags.Zero)
		} else {
			assert.False(t, fixture.flags.Zero)
		}

		assert.Equal(t, CalculateParity(uint8(i)+uint8(i)), CalculateParity(fixture.regs.A))

		if i+i < 0x80 {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}
	}
}
