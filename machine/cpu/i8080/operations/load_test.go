package operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadRegisterPairImmediate(t *testing.T) {
	for i := uint8(0); i < 4; i++ {
		pc = 0x0000
		memory[0x0000] = 0xaa
		memory[0x0001] = 0x55
		LoadRegisterPairImmediate(i)
		assert.Equal(t, uint16(0x55aa), GetRegisterPair(i))
	}
}

func TestLoadAccumulator(t *testing.T) {
	pc = 0x0000
	memory[0x0000] = 0x02
	memory[0x0001] = 0x00
	memory[0x0002] = 0xaa

	LoadAccumulator()
	assert.Equal(t, uint16(0xaa), registers[A])
}

func TestLoadHLDirect(t *testing.T) {
	pc = 0x0000
	memory[0x0000] = 0x02
	memory[0x0001] = 0x00
	memory[0x0002] = 0xaa
	memory[0x0003] = 0x55

	LoadHLDirect()
	assert.Equal(t, uint16(0x55), registers[H])
	assert.Equal(t, uint16(0xaa), registers[L])
}

func TestLoadAccumulatorIndirect(t *testing.T) {
	reset()
	memory[0x0000] = 0xaa

	LoadAccumulatorIndirect(BC)
	assert.Equal(t, uint16(0xaa), registers[A])

	memory[0x0001] = 0x55
	registers[E] = 0x0001
	LoadAccumulatorIndirect(DE)
	assert.Equal(t, uint16(0x55), registers[A])
}
