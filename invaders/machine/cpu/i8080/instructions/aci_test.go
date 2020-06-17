package instructions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// A <- A +  byte 2 (PC) + CY
func TestACI_NoCarry(t *testing.T) {
	fixture := MakeTestFixture()

	fixture.regs.PC = 0x00
	fixture.flags.Carry = false
	fixture.memory.Storage[0x00] = 0x00
	ACI(0xce, fixture.memory, fixture.regs, fixture.flags)

	assert.Equal(t, uint8(0x00), fixture.regs.A)
	assert.Equal(t, uint16(0x01), fixture.regs.PC)
	assert.True(t, fixture.flags.Zero)
	assert.False(t, fixture.flags.Sign)
	assert.False(t, fixture.flags.Carry)
	assert.True(t, fixture.flags.Parity)
	assert.False(t, fixture.flags.HalfCarry)
}

func TestACI_WithCarry(t *testing.T) {
	fixture := MakeTestFixture()

	fixture.regs.A = 0xaf
	fixture.regs.PC = 0x00
	fixture.flags.Carry = true
	fixture.memory.Storage[0x00] = 0x00
	ACI(0xce, fixture.memory, fixture.regs, fixture.flags)

	assert.Equal(t, uint8(0xb0), fixture.regs.A)
	assert.Equal(t, uint16(0x01), fixture.regs.PC)
	assert.True(t, fixture.flags.Sign)
	assert.False(t, fixture.flags.Zero)
	assert.False(t, fixture.flags.Carry)
	assert.False(t, fixture.flags.Parity)
	assert.True(t, fixture.flags.HalfCarry)
}
