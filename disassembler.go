package main

import (
	"fmt"
	"goinvaders/machine"
	"goinvaders/machine/cpu/i8080"
	"goinvaders/machine/cpu/i8080/instructions"
	"goinvaders/machine/memory"
	"os"
)

func main() {
	mem := new(memory.Memory)
	mem.LoadRom(machine.RomFilename)

	disassemblyFile, err := os.Create(machine.DisassemblyFilename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = disassemblyFile.Close() }()

	for idx := uint16(0); idx < memory.ROM_SIZE; idx += 0 {
		opcode := mem.ReadByte(idx)
		line := fmt.Sprintf("%04x\t%02x\t%s", idx, opcode, i8080.DisassembleTable[opcode])

		if instructions.OpcodesLength[opcode] == 2 {
			line = fmt.Sprintf("%s%02x", line, mem.ReadByte(idx+1))
		} else if instructions.OpcodesLength[opcode] == 3 {
			line = fmt.Sprintf("%s%02x", line, mem.ReadWord(idx+1))
		}

		_, _ = disassemblyFile.WriteString(fmt.Sprintf("%s\n", line))
		idx += uint16(instructions.OpcodesLength[opcode])
	}
}
