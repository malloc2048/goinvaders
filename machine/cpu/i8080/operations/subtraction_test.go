package operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtract(t *testing.T) {
	reset()

	for i := uint8(0); i < 6; i++ {
		registers[A] = 0x0001
		registers[i] = uint16(0x0001)
		Subtract(i)

		assert.Equal(t, uint16(0), registers[A])
		assert.True(t, Zero())
		assert.False(t, Sign())
		assert.True(t, Parity())
	}

	registers[A] = 0x0001
	Subtract(A)
	assert.Equal(t, uint16(0x0000), registers[A])

	registers[A] = 0x0001
	SetRegisterPair(HL, 0x0000)
	WriteMemoryAtHL(0xaa)
	Subtract(M)
	assert.Equal(t, uint16(0x0057), registers[A])
	assert.True(t, Carry())
}
