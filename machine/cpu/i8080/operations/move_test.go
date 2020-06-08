package operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveRegister_B(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		registers[i] = uint16(i)
		MoveRegister(i, 0x00)
		assert.Equal(t, registers[i], registers[0x00])
	}
}

func TestMoveRegister_C(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		registers[i] = uint16(i)
		MoveRegister(i, 0x01)
		assert.Equal(t, registers[i], registers[0x01])
	}
}

func TestMoveRegister_D(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		registers[i] = uint16(i)
		MoveRegister(i, 0x02)
		assert.Equal(t, registers[i], registers[0x02])
	}
}

func TestMoveRegister_E(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		registers[i] = uint16(i)
		MoveRegister(i, 0x03)
		assert.Equal(t, registers[i], registers[0x03])
	}
}

func TestMoveRegister_H(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		if i == 0x06 {
			registers[0x04] = 0x04
		}

		registers[i] = uint16(i)
		MoveRegister(i, 0x04)
		assert.Equal(t, registers[i], registers[0x04])
	}
}

func TestMoveRegister_L(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		registers[i] = uint16(i)
		MoveRegister(i, 0x05)
		assert.Equal(t, registers[i], registers[0x05])
	}
}

func TestMoveRegister_M(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		registers[i] = uint16(i)
	}
	for i := uint8(0); i < 8; i++ {
		MoveRegister(i, 0x06)

		if i != 0x06 {
			assert.Equal(t, registers[i], uint16(memory[0x0405]))
		}
	}
}

func TestMoveRegister_A(t *testing.T) {
	memory[0x0405] = 0x06
	for i := uint8(0); i < 8; i++ {
		registers[i] = uint16(i)
		MoveRegister(i, 0x07)
		assert.Equal(t, registers[i], registers[0x07])
	}
}

func TestMoveImmediate(t *testing.T) {
	reset()
	for i := uint8(0); i < 8; i++ {
		memory[uint16(i)] = 0xaa
		MoveImmediate(i)

		if i == 0x06 {
			assert.Equal(t, ReadMemoryAtHL(), memory[uint16(i)])
		} else {
			assert.Equal(t, registers[i], uint16(memory[uint16(i)]))
		}
	}
}
