package operations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreAccumulatorDirect(t *testing.T) {
	pc = 0x0000
	memory[0x0000] = 0x0002
	memory[0x0001] = 0x0000
	memory[0x0002] = 0x0000

	StoreAccumulatorDirect()
	assert.Equal(t, uint8(0xaa), memory[0x0002])
}

func TestStoreHLDirect(t *testing.T) {
	pc = 0x0000
	memory[0x0000] = 0x0002
	memory[0x0001] = 0x0000
	memory[0x0002] = 0x0000
	memory[0x0003] = 0x0000

	StoreHLDirect()
	assert.Equal(t, uint8(registers[0x05]), memory[0x0002])
	assert.Equal(t, uint8(registers[0x04]), memory[0x0003])
}

func TestStoreAccumulatorIndirect(t *testing.T) {
	reset()
	registers[A] = 0xaa
	SetRegisterPair(BC, 0x0001)
	StoreAccumulatorIndirect(BC)
	assert.Equal(t, uint16(memory[0x0001]), registers[A])

	memory[0x0001] = 0x55
	SetRegisterPair(DE, 0x0001)
	StoreAccumulatorIndirect(DE)
	assert.Equal(t, uint16(memory[0x0001]), registers[A])
}

func TestExchangeHLWithDE(t *testing.T) {
	reset()
	SetRegisterPair(DE, 0xaaaa)
	SetRegisterPair(HL, 0x5555)

	ExchangeHLWithDE()

	assert.Equal(t, uint16(0x5555), GetRegisterPair(DE))
	assert.Equal(t, uint16(0xaaaa), GetRegisterPair(HL))
}
