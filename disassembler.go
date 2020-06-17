package main

import (
	"fmt"
	"goinvaders/invaders/machine"
	"goinvaders/invaders/machine/cpu/i8080"
	"goinvaders/invaders/machine/cpu/i8080/instructions"
	"goinvaders/invaders/machine/memory"
	"os"
	"sort"
)

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	var usedOpcodes []string
	mem := new(memory.Memory)
	mem.LoadRom(machine.RomFilename)

	disassemblyFile, err := os.Create(machine.DisassemblyFilename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = disassemblyFile.Close() }()

	for idx := uint16(0); idx < memory.ROM_SIZE; idx += 0 {
		opcode := mem.Read(idx)
		usedOpcodes = append(usedOpcodes, i8080.DisassembleTable[opcode])

		line := fmt.Sprintf("%04x\t%02x\t%s", idx, opcode, i8080.DisassembleTable[opcode])

		if instructions.OpcodesLength[opcode] == 2 {
			line = fmt.Sprintf("%s%02x", line, mem.Read(idx+1))
		} else if instructions.OpcodesLength[opcode] == 3 {
			line = fmt.Sprintf("%s%02x", line, mem.ReadWord(idx+1))
		}

		_, _ = disassemblyFile.WriteString(fmt.Sprintf("%s\n", line))
		idx += uint16(instructions.OpcodesLength[opcode])
	}
	uniqueCodes := unique(usedOpcodes)
	sort.Strings(uniqueCodes)

	for _, entry := range uniqueCodes {
		fmt.Println(entry)
	}
}
