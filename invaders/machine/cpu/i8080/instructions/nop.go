package instructions

import (
	"goinvaders/invaders/machine/cpu/i8080/registers"
	"log"
)

func NOP(opcode byte, registers *registers.Registers, mnemonic string) {
	log.Printf("NOP address: %04x opcode: %02x %s\n", registers.PC, opcode, mnemonic)
	registers.PC += uint16(OpcodesLength[opcode] - 1)
}
