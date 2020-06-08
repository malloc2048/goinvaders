package operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	reset()

	for i := uint8(0); i < 6; i++ {
		registers[A] = 0x0001
		registers[i] = uint16(i)

		Add(i)
		assert.Equal(t, registers[i]+1, registers[A])
		assert.False(t, Zero())
		assert.False(t, Sign())
	}

	registers[A] = 0x0001
	Add(A)
	assert.Equal(t, uint16(0x0002), registers[A])

	SetRegisterPair(HL, 0x0001)
	WriteMemoryAtHL(0xaa)
	Add(M)
	assert.Equal(t, uint16(0x00ac), registers[A])
}

func TestAddImmediate(t *testing.T) {
	reset()
	registers[A] = 0x81
	memory[0x0000] = 0x0001

	AddImmediate()
	assert.Equal(t, uint16(0x0082), registers[A])
	assert.False(t, Zero())
	assert.True(t, Sign())
}

func TestAddWithCarry(t *testing.T) {
	reset()

	for i := uint8(0); i < 6; i++ {
		SetCarry(true)
		registers[A] = 0x01
		registers[i] = 0x01

		AddWithCarry(i)
		assert.Equal(t, uint16(0x03), registers[A])
		assert.False(t, Carry())
	}

	SetCarry(true)
	registers[A] = 0x01
	AddWithCarry(A)
	assert.Equal(t, uint16(0x0003), registers[A])

	SetRegisterPair(HL, 0x0000)
	WriteMemoryAtHL(0xaa)
	SetCarry(true)
	registers[A] = 0x01
	AddWithCarry(M)
	assert.Equal(t, uint16(0x00ac), registers[A])
}

func TestAddImmediateWithCarry(t *testing.T) {
	reset()
	registers[A] = 0x81
	SetCarry(true)
	memory[0x0000] = 0x0001

	AddImmediateWithCarry()
	assert.Equal(t, uint16(0x0083), registers[A])
}
