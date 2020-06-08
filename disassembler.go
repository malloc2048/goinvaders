package main

import (
	"fmt"
	"goinvaders/machine/cpu/i8080"
	"goinvaders/machine/memory"
	"os"
)

func main() {
	//var line string
	//var disassembly []string
	mem := new(memory.Memory)
	mem.LoadRom("roms/invaders")

	disassemblyFile, err := os.Create("roms/invaders.dis")
	if err != nil {
		panic(err)
	}
	defer func() { _ = disassemblyFile.Close() }()

	for idx := uint16(0); idx < memory.ROM_SIZE; idx += 0 {
		opcode := mem.ReadByte(idx)
		line := fmt.Sprintf("%04x\t%02x\t%s", idx, opcode, i8080.DISASSEMBLE_TABLE[opcode])

		if i8080.OPCODES_LENGTH[opcode] == 2 {
			line = fmt.Sprintf("%s%02x", line, mem.ReadByte(idx+1))
		} else if i8080.OPCODES_LENGTH[opcode] == 3 {
			line = fmt.Sprintf("%s%02x", line, mem.ReadWord(idx+1))
		}

		_, _ = disassemblyFile.WriteString(fmt.Sprintf("%s\n", line))
		idx += uint16(i8080.OPCODES_LENGTH[opcode])
	}
}
