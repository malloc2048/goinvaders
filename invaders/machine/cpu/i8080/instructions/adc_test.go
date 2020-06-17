package instructions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// A <- A + REG + CY
func TestADC_B(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.regs.B = uint8(i)
		fixture.flags.Carry = true
		ADC(0x88, fixture.memory, fixture.regs, fixture.flags)

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
	}
}

func TestADC_C(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.regs.C = uint8(i)
		fixture.flags.Carry = true
		ADC(0x89, fixture.memory, fixture.regs, fixture.flags)

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
	}
}

func TestADC_D(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.regs.D = uint8(i)
		fixture.flags.Carry = true
		ADC(0x8a, fixture.memory, fixture.regs, fixture.flags)

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
	}
}

func TestADC_E(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.regs.E = uint8(i)
		fixture.flags.Carry = true
		ADC(0x8b, fixture.memory, fixture.regs, fixture.flags)

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
	}
}

func TestADC_H(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.regs.H = uint8(i)
		fixture.flags.Carry = true
		ADC(0x8c, fixture.memory, fixture.regs, fixture.flags)

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
	}
}

func TestADC_L(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.regs.L = uint8(i)
		fixture.flags.Carry = true
		ADC(0x8d, fixture.memory, fixture.regs, fixture.flags)

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
	}
}

func TestADC_M(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = 0x00
		fixture.memory.Storage[0x0000] = uint8(i)
		ADC(0x8e, fixture.memory, fixture.regs, fixture.flags)

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

func TestADC_A(t *testing.T) {
	fixture := MakeTestFixture()

	for i := 0x00; i < 0x0100; i++ {
		fixture.regs.A = uint8(i)
		fixture.flags.Carry = true
		ADC(0x8f, fixture.memory, fixture.regs, fixture.flags)

		assert.Equal(t, uint8(i)+uint8(i)+0x01, fixture.regs.A)
		assert.Equal(t, uint16(0), fixture.regs.PC)
		assert.False(t, fixture.flags.Zero)

		assert.Equal(t, CalculateParity(uint8(i)+uint8(i)+0x01), CalculateParity(fixture.regs.A))

		if i+i+0x01 < 0x80 {
			assert.False(t, fixture.flags.Sign)
		} else {
			assert.True(t, fixture.flags.Sign)
		}
	}
}
