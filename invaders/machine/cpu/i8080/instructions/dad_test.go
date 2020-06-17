package instructions

import (
	"github.com/stretchr/testify/assert"
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"testing"
)

func TestDAD_B(t *testing.T) {
	regs := new(registers.Registers)
	flags := new(registers.Flags)
	opcode := byte(0x09)

	flags.Carry = false
	regs.B = 0x22
	regs.C = 0x00
	regs.H = 0x22
	regs.L = 0x22

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0x44), regs.H)
	assert.Equal(t, uint8(0x22), regs.L)
	assert.False(t, flags.Carry)

	regs.B = 0xff
	regs.C = 0xff

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0x44), regs.H)
	assert.Equal(t, uint8(0x21), regs.L)
	assert.True(t, flags.Carry)
}

func TestDAD_D(t *testing.T) {
	regs := new(registers.Registers)
	flags := new(registers.Flags)
	opcode := byte(0x19)

	flags.Carry = false
	regs.D = 0x22
	regs.E = 0x00
	regs.H = 0x22
	regs.L = 0x22

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0x44), regs.H)
	assert.Equal(t, uint8(0x22), regs.L)
	assert.False(t, flags.Carry)

	regs.D = 0xff
	regs.E = 0xff

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0x44), regs.H)
	assert.Equal(t, uint8(0x21), regs.L)
	assert.True(t, flags.Carry)
}

func TestDAD_H(t *testing.T) {
	regs := new(registers.Registers)
	flags := new(registers.Flags)
	opcode := byte(0x29)

	flags.Carry = false
	regs.H = 0x22
	regs.L = 0x22

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0x44), regs.H)
	assert.Equal(t, uint8(0x44), regs.L)
	assert.False(t, flags.Carry)

	regs.H = 0xff
	regs.L = 0xff

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0xff), regs.H)
	assert.Equal(t, uint8(0xfe), regs.L)
	assert.True(t, flags.Carry)
}

func TestDAD_SP(t *testing.T) {
	regs := new(registers.Registers)
	flags := new(registers.Flags)
	opcode := byte(0x39)

	flags.Carry = false
	regs.SP = 0x2200
	regs.H = 0x22
	regs.L = 0x22

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0x44), regs.H)
	assert.Equal(t, uint8(0x22), regs.L)
	assert.False(t, flags.Carry)

	regs.SP = 0xffff

	DAD(opcode, regs, flags)
	assert.Equal(t, uint8(0x44), regs.H)
	assert.Equal(t, uint8(0x21), regs.L)
	assert.True(t, flags.Carry)
}
